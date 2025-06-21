import { useState, useEffect, useRef } from "react";
import { createReceiptComplete } from "../api/receipts";
import type { ReceiptCreateComplete } from "../types/receipt";
import type { StoreProductWithDetails } from "../types/store_product";
import type { Product } from "../types/product";
import type { CustomerCard } from "../types/customer_card";
import { fetchStoreProductsWithDetails } from "../api/store_products";
import { fetchProducts } from "../api/products";
import { getCustomerCards } from "../api/customer_cards";
import { useAuth } from "../contexts/AuthContext";

const CreateReceipt = () => {
  const { user } = useAuth();
  const [cardNumber, setCardNumber] = useState<string | undefined>(undefined);
  const [items, setItems] = useState<{ upc: string; product_number: number }[]>(
    [],
  );
  const [products, setProducts] = useState<StoreProductWithDetails[]>([]);
  const [allProducts, setAllProducts] = useState<Product[]>([]);
  const [customerCards, setCustomerCards] = useState<CustomerCard[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<{ receipt_number: string } | null>(
    null,
  );
  const itemRefs = useRef<(HTMLSelectElement | null)[]>([]);

  useEffect(() => {
    fetchStoreProductsWithDetails().then((res) => setProducts(res.data || []));
    fetchProducts().then((res) => setAllProducts(res.data || []));
    getCustomerCards().then(setCustomerCards);
  }, []);

  const addItem = () => {
    if (items.length > 0) {
      const last = items[items.length - 1];
      if (!last.upc || !last.product_number) return;
    }
    setItems([...items, { upc: "", product_number: 1 }]);
    setTimeout(() => {
      if (itemRefs.current[items.length])
        itemRefs.current[items.length]?.focus();
    }, 100);
  };

  const updateItem = (
    idx: number,
    field: "upc" | "product_number",
    value: string | number,
  ) => {
    setItems(
      items.map((item, i) => (i === idx ? { ...item, [field]: value } : item)),
    );
  };

  const removeItem = (idx: number) => {
    setItems(items.filter((_, i) => i !== idx));
  };

  // Calculate price for an item
  const getItemPrice = (upc: string, quantity: number) => {
    const product = products.find((p) => p.upc === upc);
    if (!product) return { price: 0, isPromo: false, total: 0 };
    let price = product.selling_price;
    const isPromo = product.promotional_product;
    if (isPromo) price = price * 0.8;
    return { price, isPromo, total: price * quantity };
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(null);

    if (!user?.employee_id) {
      setError("User authentication error. Please log in again.");
      setLoading(false);
      return;
    }

    if (items.length === 0) {
      setError("Please add at least one item.");
      setLoading(false);
      return;
    }
    if (
      items.some((i) => !i.upc || !i.product_number || i.product_number < 1)
    ) {
      setError("Please select a product and quantity for each item.");
      setLoading(false);
      return;
    }
    try {
      const now = new Date();
      const print_date = now.toISOString().slice(0, 19).replace("T", " ");
      const receipt: ReceiptCreateComplete = {
        employee_id: user.employee_id,
        card_number: cardNumber || null,
        print_date,
        items: items.map((i) => {
          const { price } = getItemPrice(i.upc, i.product_number);
          return {
            upc: i.upc,
            product_number: parseInt(i.product_number.toString()),
            selling_price: price,
          };
        }),
      };
      console.log("Sending receipt data:", JSON.stringify(receipt, null, 2));
      const res = await createReceiptComplete(receipt);
      setSuccess({ receipt_number: res.data.id });
      setItems([]);
      setCardNumber(undefined);
    } catch (err: unknown) {
      console.error("Receipt creation error:", err);
      console.error("Error response:", (err as any)?.response);
      const errorMessage =
        (err as any)?.response?.data?.error ||
        (err as any)?.response?.data?.message ||
        (err as any)?.message ||
        "Failed to create receipt";
      setError(`Error: ${errorMessage}`);
    } finally {
      setLoading(false);
    }
  };

  // Calculate total for all items
  const subtotal = items.reduce((sum, i) => {
    const { total } = getItemPrice(i.upc, i.product_number);
    return sum + total;
  }, 0);

  // Get discount percent from selected card
  const selectedCard = customerCards.find(
    (card) => card.card_number === cardNumber,
  );
  const discountPercent = selectedCard ? selectedCard.percent : 0;
  const discountAmount = subtotal * (discountPercent / 100);
  const total = subtotal - discountAmount;

  return (
    <div className="max-w-2xl mx-auto p-6 bg-white rounded shadow">
      <h1 className="text-2xl font-bold mb-4">Create Receipt</h1>
      {error && (
        <div className="bg-red-100 text-red-700 p-2 mb-2 rounded">{error}</div>
      )}
      {success && (
        <div className="bg-green-100 text-green-700 p-2 mb-2 rounded flex flex-col gap-2">
          <div>
            Receipt created! Receipt Number: <b>{success.receipt_number}</b>
          </div>
          <button
            className="bg-blue-500 text-white px-3 py-1 rounded w-fit"
            onClick={() => setSuccess(null)}
          >
            Create another receipt
          </button>
        </div>
      )}
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="font-medium">Cashier</label>
          <div className="w-full border p-2 rounded mt-1 bg-gray-50">
            {user ? (
              <div className="font-medium">
                {user.empl_surname} {user.empl_name}{" "}
                {user.empl_patronymic || ""} ({user.employee_id})
              </div>
            ) : (
              <div className="text-gray-500">Loading user information...</div>
            )}
          </div>
          <div className="text-xs text-gray-500 mt-1">
            Receipt will be created under your account.
          </div>
        </div>
        <div>
          <label className="font-medium">Customer Card (optional)</label>
          <select
            value={cardNumber || ""}
            onChange={(e) => setCardNumber(e.target.value || undefined)}
            className="w-full border p-2 rounded mt-1"
          >
            <option value="">No card</option>
            {customerCards.map((card) => (
              <option key={card.card_number} value={card.card_number}>
                {card.cust_surname} {card.cust_name}{" "}
                {card.cust_patronymic || ""} ({card.card_number}) —{" "}
                {card.percent}%
              </option>
            ))}
          </select>
          <div className="text-xs text-gray-500 mt-1">
            If the customer has a loyalty card, select it to apply a discount.
          </div>
        </div>
        <div>
          <label className="font-medium">Items</label>
          <button
            type="button"
            onClick={addItem}
            className="ml-2 px-2 py-1 bg-blue-500 text-white rounded text-sm"
          >
            Add Item
          </button>
          <div className="overflow-x-auto mt-2">
            <table className="min-w-full border text-sm">
              <thead>
                <tr className="bg-gray-100">
                  <th className="px-2 py-1 border">Product</th>
                  <th className="px-2 py-1 border">Qty</th>
                  <th className="px-2 py-1 border">Price</th>
                  <th className="px-2 py-1 border">Total</th>
                  <th className="px-2 py-1 border">Remove</th>
                </tr>
              </thead>
              <tbody>
                {items.map((item, idx) => {
                  const { price, isPromo, total } = getItemPrice(
                    item.upc,
                    item.product_number,
                  );
                  const product = products.find((p) => p.upc === item.upc);
                  const characteristics = product
                    ? allProducts.find(
                        (prod) => prod.product_id === product.product_id,
                      )?.characteristics
                    : undefined;
                  return (
                    <tr key={idx}>
                      <td className="border px-2 py-1">
                        <select
                          ref={(el) => {
                            itemRefs.current[idx] = el;
                          }}
                          value={item.upc}
                          onChange={(e) =>
                            updateItem(idx, "upc", e.target.value)
                          }
                          className="border p-1 rounded w-full"
                        >
                          <option value="">Select product</option>
                          {products.map((p) => (
                            <option key={p.upc} value={p.upc}>
                              {p.product_name} ({p.upc})
                            </option>
                          ))}
                        </select>
                        {characteristics && (
                          <div className="text-xs text-gray-500 mt-1">
                            {characteristics}
                          </div>
                        )}
                      </td>
                      <td className="border px-2 py-1">
                        <input
                          type="number"
                          min={1}
                          value={item.product_number}
                          onChange={(e) =>
                            updateItem(idx, "product_number", e.target.value)
                          }
                          className="w-16 border p-1 rounded"
                        />
                      </td>
                      <td className="border px-2 py-1 text-right">
                        {item.upc &&
                          product &&
                          (isPromo ? (
                            <span className="text-orange-600 font-semibold">
                              ${price.toFixed(2)}{" "}
                              <span className="text-xs">(Promo -20%)</span>
                            </span>
                          ) : (
                            <span>${price.toFixed(2)}</span>
                          ))}
                      </td>
                      <td className="border px-2 py-1 text-right">
                        {item.upc ? `$${total.toFixed(2)}` : ""}
                      </td>
                      <td className="border px-2 py-1 text-center">
                        <button
                          type="button"
                          onClick={() => removeItem(idx)}
                          className="text-red-500 hover:underline"
                        >
                          Remove
                        </button>
                      </td>
                    </tr>
                  );
                })}
                {items.length === 0 && (
                  <tr>
                    <td colSpan={5} className="text-center text-gray-400 py-4">
                      No items added.
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        </div>
        <div className="text-right font-bold text-lg mt-4">
          Subtotal: ${subtotal.toFixed(2)}
          <br />
          {discountPercent > 0 && (
            <span className="text-green-700">
              Discount ({discountPercent}%): -${discountAmount.toFixed(2)}
            </span>
          )}
          <br />
          Final Total: ${total.toFixed(2)}
        </div>
        <button
          type="submit"
          className="bg-green-600 text-white px-4 py-2 rounded mt-2"
          disabled={loading}
        >
          {loading ? "Creating..." : "Create Receipt"}
        </button>
      </form>
    </div>
  );
};

export default CreateReceipt;
