import { defineConfig } from "orval";

export default defineConfig({
  bbs: {
    input: "../tsp-output/schema/openapi.yaml",
    output: {
      mode: "single",
      target: "./src/generated/api.ts",
      schemas: "./src/generated/model",
      client: "react-query",
      baseUrl: "/api",
      override: {
        mutator: {
          path: "./src/api/fetcher.ts",
          name: "customFetch",
        },
      },
    },
  },
});
