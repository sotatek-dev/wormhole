import Caver from 'caver-js';
import klaytnBridgeImplementationAbi from '../abi/BridgeImplementation.json';

const TOKEN_BRIDGE_ADDRESS_KLAYTN = process.env.REACT_APP_TOKEN_BRIDGE;
const KLAYTN_PROVIDER_API = process.env.REACT_APP_KLAYTN_PROVIDER_API;
const KEY_KLAYTN = process.env.REACT_APP_KEY_KLAYTN;

const GAS_PRICE_KLAYTN = "25000000000";
export default async function isWrappedAsset(
    address?: string, 
) {
    
    const caver = new Caver(KLAYTN_PROVIDER_API);

    const contract = new caver.klay.Contract(klaytnBridgeImplementationAbi as any, TOKEN_BRIDGE_ADDRESS_KLAYTN, {
        from: KEY_KLAYTN,
        gasPrice: GAS_PRICE_KLAYTN
    })

    const result = await contract.methods.isWrappedAsset(address).call();
    
    return result;
}