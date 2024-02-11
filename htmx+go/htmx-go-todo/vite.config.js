import { resolve } from "path";
import { defineConfig } from "vite";

export default defineConfig({
	build: {
		lib: {
			entry: resolve(__dirname, "index.js"),
			formats: ["es"],
			name: "[name]",
			filename: "[name]",
		},
		outDir: "assets/js",
		emptyOutDir: false,
	},
});
