# Solidity API

## IVoucher

Interface describing Voucher for redeeming game items

\_The Voucher type must have a strict implementation on backend

A Voucher is a document signed from the server IssuerKey and allows the execution
of actions on the game generally for creating game items, such as Booster Packs, Meld or reward tokens\_

### Voucher

```solidity
struct Voucher {
  address target;
  uint8 action;
  address beneficiary;
  address referrer;
  uint64 expire;
  uint32 chainId;
  bytes32 voucherCodeHash;
  bytes encodedParams;
}
```

### useVouchers

```solidity
function useVouchers(struct IVoucher.Voucher[] vouchers, bytes signature) external
```

Use vouchers that were issued and signed by the Back-end to receive game items.

#### Parameters

| Name      | Type                      | Description                      |
| --------- | ------------------------- | -------------------------------- |
| vouchers  | struct IVoucher.Voucher[] | Vouchers issued by the Back-end. |
| signature | bytes                     | Vouchers signed by the Back-end. |

### useVoucher

```solidity
function useVoucher(struct IVoucher.Voucher voucher, bytes signature) external
```

Use the voucher that was signed by the Back-end to receive game items.

#### Parameters

| Name      | Type                    | Description                     |
| --------- | ----------------------- | ------------------------------- |
| voucher   | struct IVoucher.Voucher | Voucher issued by the Back-end. |
| signature | bytes                   | Voucher signed by the Back-end. |
