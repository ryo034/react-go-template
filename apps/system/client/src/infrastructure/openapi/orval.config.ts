export default {
  system: {
    input: "../../../../../../schema/api/system/openapi/dist/openapi.yaml",
    output: {
      target: "../../generated/schema/openapi/systemApi.ts",
      client: "zod"
    }
  }
}
