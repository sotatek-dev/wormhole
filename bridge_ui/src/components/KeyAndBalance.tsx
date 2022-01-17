import {
  ChainId,
  CHAIN_ID_SOLANA,
  CHAIN_ID_TERRA,
  CHAIN_ID_KLAYTN_BAOBAB,
  isEVMChain,
} from "@certusone/wormhole-sdk";
import EthereumSignerKey from "./EthereumSignerKey";
import KaikasWalletKey from "./KlaytnWalletKey";
import SolanaWalletKey from "./SolanaWalletKey";
import TerraWalletKey from "./TerraWalletKey";

function KeyAndBalance({ chainId }: { chainId: ChainId }) {
  if (isEVMChain(chainId)) {
    return (
      <>
        <EthereumSignerKey />
      </>
    );
  }
  if (chainId === CHAIN_ID_SOLANA) {
    return (
      <>
        <SolanaWalletKey />
      </>
    );
  }
  if (chainId === CHAIN_ID_TERRA) {
    return (
      <>
        <TerraWalletKey />
      </>
    );
  }
  if (chainId === CHAIN_ID_KLAYTN_BAOBAB) {
    return (
      <>
        <KaikasWalletKey />
      </>
    );
  }
  return null;
}

export default KeyAndBalance;
