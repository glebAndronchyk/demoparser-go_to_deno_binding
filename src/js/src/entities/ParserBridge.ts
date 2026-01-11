// interface BridgeScope {
// }

import { GoCnstrct } from "../types/GoCnstrct.ts";

// !!!IMPORTANT. CONTAINER FOR THE GOLANG PARSER
global.__PARSER_API__ = {} // todo: should be a pool of parsers consider using approach below
// https://github.com/golang/go/issues/25612#issuecomment-491919708
export class ParserBridge extends EventTarget {
    private api: never;

    // todo: do not create GoCnstrct twice
    private static gojs = new (globalThis as any).Go() as GoCnstrct;

    static async construct() {
        const binaryUrl = new URL("../../wasm/main.wasm", import.meta.url);
        const binary = await Deno.readFile(binaryUrl);
        const wasm = await WebAssembly.instantiate(binary, this.gojs.importObject);
        this.gojs.run(wasm.instance);

        this.create();

        // todo: properly handle kill of the go process
        Deno.addEventListener("beforeunload", this.kill)

        return new this();
    }

    private constructor() {
        super()
        this.api = global.__PARSER_API__
    }

    private static kill() {
        global.__PARSER_API__.Kill();
    }

    private static create() {
        global.__PARSER_API__.Create(Deno.readFileSync("./match.dem"));
        Object.freeze(global.__PARSER_API__)
    }

    parseToEnd() {
        this.api.ParseToEnd()
    }

    parseNextFrame() {
        this.api.ParseNextFrame()
    }

    getEntityState(handle: number) {
        return this.api.GetEntityState(handle);
    }

    getGameState() {
        return this.api.GetStaticGameState();
    }

    addEventListener(type: string, callback: Function): number {
        return this.api.RegisterEvent(type, {
            onExecute: callback,
        });
    }

    override removeEventListener(handle: number) {
        this.api.UnregisterEvent(handle);
    }
}
