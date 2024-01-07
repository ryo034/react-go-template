export interface HttpClient {
  get<T>(url: string, params?: any): Promise<T>
}
