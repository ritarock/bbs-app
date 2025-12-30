import { serve } from "bun";
import index from "./index.html";

const API_SERVER = "http://localhost:8080";

const server = serve({
  routes: {
    // API Proxy: forward /api/* to backend server
    "/api/*": async (req) => {
      const url = new URL(req.url);
      // Remove /api prefix and forward to backend
      const targetPath = url.pathname.replace(/^\/api/, "");
      const targetUrl = `${API_SERVER}${targetPath}${url.search}`;

      const headers = new Headers(req.headers);
      headers.delete("host");

      const proxyReq = new Request(targetUrl, {
        method: req.method,
        headers,
        body: req.body,
      });

      try {
        return await fetch(proxyReq);
      } catch (error) {
        console.error("Proxy error:", error);
        return Response.json(
          { code: 502, message: "Backend server unavailable" },
          { status: 502 }
        );
      }
    },

    // Serve index.html for all unmatched routes.
    "/*": index,
  },

  development: process.env.NODE_ENV !== "production" && {
    // Enable browser hot reloading in development
    hmr: true,

    // Echo console logs from the browser to the server
    console: true,
  },
});

console.log(`🚀 Server running at ${server.url}`);
