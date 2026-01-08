import '../wasm/wasm_exec.js';
import { ParserBridge } from './entities/ParserBridge.ts';

const bridge = await ParserBridge.construct();
const gs = await bridge.getGameState();

console.log(gs);


export default ParserBridge;
