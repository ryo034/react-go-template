import redis, { Redis } from "ioredis"

interface InMemoryClient {
  get(key: string): Promise<string | null>
}

export class RedisClient implements InMemoryClient {
  constructor(readonly client: Redis) {}

  static async create() {
    const client = await redis.createClient()
    return new RedisClient(client)
  }

  async get(key: string): Promise<string | null> {
    return this.client.get(key)
  }
}
