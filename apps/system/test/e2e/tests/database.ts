import { readFileSync } from "fs"
import { parse } from "csv-parse/sync"
import { Client, types } from "pg"
import format from "pg-format"
import { dbConfig, env, initializeData } from "./config"

interface Database {
  getConnection(): Promise<Client>
  canConnect(): Promise<boolean>
  setup(): Promise<void>
  clear(): Promise<void>
}

export const targetTables = [
  "address_components",
  "system_accounts",
  "system_account_profiles",
  "system_account_phone_numbers",
  "workspaces",
  "workspace_details",
  "members",
  "member_profiles",
  "member_addresses",
  "membership_periods"
]

export class MainDb implements Database {
  async getConnection() {
    types.setTypeParser(1114, (stringValue) => stringValue)
    const client = new Client(dbConfig)
    await client.connect()
    return client
  }

  async canConnect(): Promise<boolean> {
    try {
      await this.getConnection()
      return true
    } catch (_e) {
      return false
    }
  }

  async setup(): Promise<void> {
    const connection = await this.getConnection()
    for (const tableName of targetTables) {
      await this.insertInitDataInCsvPostgres(connection, tableName)
    }

    connection.end()
  }

  private async insertInitDataInCsvPostgres(connection: Client, tableName: string) {
    let csv: Array<Array<string>> = []
    if (initializeData === "true" && env !== "localhost") {
      csv = parse(readFileSync(`./setup/database/${tableName}.csv`))
    } else {
      csv = parse(readFileSync(`./setup/database/${tableName}.csv`))
    }

    const header = csv.shift()
    if (!header) {
      throw new Error("csv header is empty")
    }
    const columns = header.join(",")

    if (csv.length > 0) {
      for (const row of csv) {
        const query = format(`INSERT INTO ${tableName} (${columns}) VALUES %L`, [row]).replace(/''/gi, "NULL")
        await connection.query(query)
      }
    }
  }

  async clear(): Promise<void> {
    const connection = await this.getConnection()
    const reversedTables = [...targetTables].reverse()
    for (const tableName of reversedTables) {
      const query = `DELETE FROM ${tableName};`
      await connection.query(query)
    }
    connection.end()
  }
}

export class Databases {
  constructor(private readonly databases: Database[]) {}

  static gen(): Databases {
    return new Databases([new MainDb()])
  }

  async setup(): Promise<void> {
    try {
      if (initializeData === "true" || env === "localhost") {
        console.log("starting databases setup....")
      } else {
        console.log("skipping databases setup....")
        return
      }
      for (const v of this.databases) {
        console.log("check canConnect...")
        const canConnect = await v.canConnect()
        console.log(`canconnect done: ${canConnect}`)
        if (canConnect) {
          await v.clear()
          await v.setup()
        }
      }
    } catch (error) {
      console.warn(`
      ==================
      Database error: ${error}
      ==================
      `)
    }
  }

  async clear(): Promise<void> {
    for (const v of this.databases) {
      if (await v.canConnect()) {
        await v.clear()
      }
    }
  }
}
