import {
  ChainId,
  CHAIN_ID_ETH,
} from "@certusone/wormhole-sdk";
import { WormholeAbi__factory } from "@certusone/wormhole-sdk/lib/esm/ethers-contracts/abi";
import { getAddress as getEthAddress } from "@ethersproject/address";
import React, { useCallback } from "react";
import { useSelector } from "react-redux";
import useIsWalletReady from "../../hooks/useIsWalletReady";
import { DataWrapper } from "../../store/helpers";
import { NFTParsedTokenAccount } from "../../store/nftSlice";
import {
  selectNFTSourceParsedTokenAccount,
  selectTransferSourceParsedTokenAccount,
} from "../../store/selectors";
import { ParsedTokenAccount } from "../../store/transferSlice";
import {
  getMigrationAssetMap,
  WORMHOLE_V1_ETH_ADDRESS,
} from "../../utils/consts";
import {
  isValidEthereumAddress,
} from "../../utils/ethereum";
import TokenPicker, { BasicAccountRender } from "./TokenPicker";
import { useKaikasProvider } from "../../contexts/KaikasProviderContext";
import { getKlaytnNFT, getKlaytnToken, klaytnNFTToNFTParsedTokenAccount, klaytnTokenToParsedTokenAccountNFT } from "../../utils/klaytn";
const isWormholev1 = (provider: any, address: string, chainId: ChainId) => {
  if (chainId !== CHAIN_ID_ETH) {
    return Promise.resolve(false);
  }
  
  const connection = WormholeAbi__factory.connect(
    WORMHOLE_V1_ETH_ADDRESS,
    provider
  );
  return connection.isWrappedAsset(address);
};

type EthereumSourceTokenSelectorProps = {
  value: ParsedTokenAccount | null;
  onChange: (newValue: ParsedTokenAccount | null) => void;
  tokenAccounts: DataWrapper<ParsedTokenAccount[]> | undefined;
  disabled: boolean;
  resetAccounts: (() => void) | undefined;
  chainId: ChainId;
  nft?: boolean;
};

export default function KlaytnTokenPicker(
  props: EthereumSourceTokenSelectorProps
) {
  const {
    value,
    onChange,
    tokenAccounts,
    disabled,
    resetAccounts,
    chainId,
    nft,
  } = props;
  const { provider: providerKaikas, signerAddress: signerAddressKaikas } = useKaikasProvider();
  const { isReady } = useIsWalletReady(chainId);
  const selectedTokenAccount: NFTParsedTokenAccount | undefined = useSelector(
    nft
      ? selectNFTSourceParsedTokenAccount
      : selectTransferSourceParsedTokenAccount
  );

  const shouldDisplayBalance = useCallback(
    (tokenAccount: NFTParsedTokenAccount) => {
      const selectedMintMatch =
        selectedTokenAccount &&
        selectedTokenAccount.mintKey.toLowerCase() ===
          tokenAccount.mintKey.toLowerCase();
      //added just in case we start displaying NFT balances again.
      const selectedTokenIdMatch =
        selectedTokenAccount &&
        selectedTokenAccount.tokenId === tokenAccount.tokenId;

      return !!(
        tokenAccount.isNativeAsset || //The native asset amount isn't taken from covalent, so can be trusted.
        (selectedMintMatch && selectedTokenIdMatch)
      );
    },
    [selectedTokenAccount]
  );

  const isMigrationEligible = useCallback(
    (address: string) => {
      const assetMap = getMigrationAssetMap(chainId);
      return !!assetMap.get(getEthAddress(address));
    },
    [chainId]
  );

  // Promise<NFTParsedTokenAccount>
  const getAddress: (
    address: string,
    tokenId?: string
  ) => Promise<NFTParsedTokenAccount> = useCallback(
    async (address: string, tokenId?: string) => {
      if (providerKaikas && signerAddressKaikas && isReady) {
        try {
          const tokenAccount = await (nft
            ? getKlaytnNFT(address, providerKaikas)
            : getKlaytnToken(address, providerKaikas));
          console.log({tokenAccount});
          if (!tokenAccount) {
            return Promise.reject("Could not find the specified token.");
          }
          if (nft && !tokenId) {
            return Promise.reject("Token ID is required.");
          } else if (nft && tokenId) {
            return klaytnNFTToNFTParsedTokenAccount(
              tokenAccount,
              tokenId,
              signerAddressKaikas
            );
          } else {
            return klaytnTokenToParsedTokenAccountNFT(
              tokenAccount,
              signerAddressKaikas,
            );
          }
        } catch (e) {
          console.error(e);
          return Promise.reject("Unable to retrive the specific token.");
        }
      } else {
        return Promise.reject({ error: "Wallet is not connected." });
      }
    },
    [isReady, nft, providerKaikas, signerAddressKaikas]
  );

  const onChangeWrapper = useCallback(
    async (account: NFTParsedTokenAccount | null) => {
      if (account === null) {
        onChange(null);
        return Promise.resolve();
      }
      let v1 = false;
      try {
          v1 = await isWormholev1(providerKaikas, account.publicKey, chainId); 
      } catch (e) {
        //For now, just swallow this one.
      }
      const migration = isMigrationEligible(account.publicKey);
      if (v1 === true && !migration) {
        throw new Error(
          "Wormhole v1 assets cannot be transferred with this bridge."
        );
      }
      onChange(account);
      return Promise.resolve();
    },
    [chainId, onChange, isMigrationEligible, providerKaikas]
  );

  const RenderComp = useCallback(
    ({ account }: { account: NFTParsedTokenAccount }) => {
      return BasicAccountRender(
        account,
        isMigrationEligible,
        nft || false,
        shouldDisplayBalance
      );
    },
    [nft, isMigrationEligible, shouldDisplayBalance]
  );

  return (
    <TokenPicker
      value={value}
      options={tokenAccounts?.data || []}
      RenderOption={RenderComp}
      useTokenId={nft}
      onChange={onChangeWrapper}
      isValidAddress={isValidEthereumAddress}
      getAddress={getAddress}
      disabled={disabled}
      resetAccounts={resetAccounts}
      error={""}
      showLoader={tokenAccounts?.isFetching}
      nft={nft || false}
      chainId={chainId}
    />
  );
}
