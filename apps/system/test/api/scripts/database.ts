import { readFileSync } from "fs"
import { parse } from "csv-parse/sync"
import { Client, types } from "pg"
import format from "pg-format"
import { dbConfig } from "./config"

export interface Database {
  canConnect(): Promise<boolean>
  setup(): Promise<void>
  clear(): Promise<void>
  getConnection(): any
}

const targetTables = [
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
    } catch (_) {
      return false
    }
  }

  async setup(): Promise<void> {
    const connection = await this.getConnection()
    for (const tableName of targetTables) {
      await this.insertInitDataInCsvPostgres(connection, tableName)
    }
    return connection.end()
  }

  private async insertInitDataInCsvPostgres(connection: Client, tableName: string) {
    let csv: Array<Array<string>> = []
    csv = parse(readFileSync(`./setup/database/${tableName}.csv`))

    const header = csv.shift()
    const columns = header.join(",")

    if (csv.length > 0) {
      for (const row of csv) {
        const query = format(`INSERT INTO ${tableName} (${columns}) VALUES %L`, [row]).replace(/''/gi, "NULL")
        console.log(`${query.substring(0, 100)}...`)
        await connection.query(query)
      }
    }
  }

  async clear(): Promise<void> {
    const connection = await this.getConnection()
    const reversedTables = [...targetTables].reverse()
    for (const tableName of reversedTables) {
      const query = `DELETE FROM ${tableName};`
      console.log(`${query.substring(0, 100)}...`)
      await connection.query(query)
    }
    return connection.end()
  }
}
