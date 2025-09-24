    export interface GoCnstrct {
        readonly importObject: WebAssembly.Imports;
        run(instance: WebAssembly.Instance): Promise<void> | void;
    }