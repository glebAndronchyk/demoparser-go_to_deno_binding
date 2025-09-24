import { emptyDir, build } from '@deno/dnt';

const BUILD_FOLDER = "./npm_build";

await emptyDir(BUILD_FOLDER);

await build({
    packageManager: 'pnpm',
    entryPoints: ['src/main.ts'],
    outDir: BUILD_FOLDER,
    scriptModule: false,
    shims: {
        deno: true,
    },
    package: {
        name: "demoparser-bindings",
        version: Deno.args[0],
        files: ["./src/wasm/main.wasm"]
    },
    compilerOptions: {
        lib: ["ESNext", "DOM"],
    },
    postBuild() {
        const nodeProjectWasmEntryPath = `${BUILD_FOLDER}/src/wasm/main.wasm`;

        console.log(`[postbuild] Copy wasm entry into ${nodeProjectWasmEntryPath}...`);
        Deno.copyFileSync('wasm/main.wasm', nodeProjectWasmEntryPath);

        console.log("[postbuild] Wait for ESM build...");
        let occurences = 0;

        const waitEsmBuildInterval = setInterval(() => {
            try {
                if (occurences > 10) {
                    console.log("[postbuild] Seems file path for wasm folder is incorrect. Breaking...");
                    clearInterval(waitEsmBuildInterval);
                    Deno.exit();
                }

               const wasmBuildFolder = `${BUILD_FOLDER}/esm/wasm1`;
               const info = Deno.lstatSync(wasmBuildFolder);
                if (info.isDirectory) {
                    console.log("[postbuild] ESM build succeded...");
                    clearInterval(waitEsmBuildInterval);
                    Deno.copyFileSync(nodeProjectWasmEntryPath, wasmBuildFolder + '/main.wasm')
                } else {
                    occurences++;
                }
            } catch (e) {
                occurences++;
            }
        }, 500);
    },
    filterDiagnostic(diagnostic) {
        if (diagnostic.file?.fileName.includes("_dnt")) {
            console.log("[filterDiagnostic] Had diagnostic issues in _dnt shims.");
            return false;
        };
        return true;
    }
});
