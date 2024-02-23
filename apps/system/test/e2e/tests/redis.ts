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

export const getOtpCodeFromRedis = async (email: string): Promise<string> => {
  const redis = await RedisClient.create()
  const code = await redis.get(`otp:${email}`)
  if (code === null) {
    return ""
  }
  return code
}
