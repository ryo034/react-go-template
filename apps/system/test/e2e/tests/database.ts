import { readFileSync } from "fs"
import { parse } from "csv-parse/sync"
import { format } from "node-pg-format"
import { Client, types } from "pg"
import { dbConfig, env, initializeData, isSilent } from "./config"

const logWithEllipsis = (text: string, maxLength = 50) => {
  if (text.length > maxLength) {
    showLog(`${text.substring(0, maxLength)}...`)
  } else {
    showLog(text)
  }
}

const showLog = (text: string) => {
  if (isSilent) return
  console.log(text)
}

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
  "member_login_histories",
  "member_profiles",
  "member_addresses",
  "membership_periods",
  "invitations"
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
    await connection.end()
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
        logWithEllipsis(query)
        await connection.query(query)
      }
    }
  }

  async clear(): Promise<void> {
    const connection = await this.getConnection()
    const reversedTables = [...targetTables].reverse()
    for (const tableName of reversedTables) {
      const query = `DELETE FROM ${tableName};`
      logWithEllipsis(query)
      await connection.query(query)
    }
    connection.end()
  }
}
