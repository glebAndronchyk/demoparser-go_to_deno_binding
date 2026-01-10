// interface BridgeScope {
// }

import { GoCnstrct } from "../types/GoCnstrct.ts";
import {AddEventListenerOptions, EventListenerOrEventListenerObject} from "npm:undici-types@7.10.0/patch.d.ts";

export class ParserBridge extends EventTarget {
    ///#region static
    private static gojs = new (globalThis as any).Go() as GoCnstrct;

    static async construct() {
        const binaryUrl = new URL("../../wasm/main.wasm", import.meta.url);
        const binary = await Deno.readFile(binaryUrl);
        const wasm = await WebAssembly.instantiate(binary, this.gojs.importObject);
        this.gojs.run(wasm.instance);
        // todo:
        global.CreateParserInstance(Deno.readFileSync("./match.dem"));

        return new this();
    }
    ///#endregion static

    // private _scope: BridgeScope;

    private constructor() {
        // this._scope = scope;
    }

    parseToEnd() {
        global.ParseToEnd()
    }

    getEntityState(handle: number) {
        const { promise, resolve } = Promise.withResolvers();

        global.GetEntityState((data) => resolve(data), handle)

        return promise;
    }

    getGameState() {
        const { promise, resolve } = Promise.withResolvers();

        global.GetGameState((data) => resolve(data))

        return promise;
    }

    addEventListener(type: string, callback: EventListenerOrEventListenerObject | null, options?: AddEventListenerOptions | boolean) {
        global.RegisterEvent("frame-done", callback);
    }
}
