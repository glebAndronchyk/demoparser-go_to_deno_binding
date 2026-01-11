import '../wasm/wasm_exec.js';
import { ParserBridge } from './entities/ParserBridge.ts';

const bridge = await ParserBridge.construct();

const handleId = bridge.addEventListener("frame-done", () => {
    console.log(bridge.getGameState());
})
// bridge.parseToEnd();

console.log({ handleId });
bridge.parseNextFrame();
bridge.removeEventListener(handleId)
bridge.parseNextFrame();

export default ParserBridge;
