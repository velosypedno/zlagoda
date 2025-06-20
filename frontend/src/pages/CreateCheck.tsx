import { useState, useEffect } from "react";
import { createCheck } from "../api/checks";
import type { CheckCreate } from "../types/check";
import type { StoreProductWithDetails } from "../types/store_product";
import type { CustomerCard } from "../types/customer_card";
import type { Employee } from "../types/employee";
import { fetchStoreProductsWithDetails } from "../api/store_products";
import { getCustomerCards } from "../api/customer_cards";
import { fetchEmployees } from "../api/employees";

const CreateCheck = () => {
  const [employeeId, setEmployeeId] = useState("");
  const [cardNumber, setCardNumber] = useState<string | undefined>(undefined);
  const [items, setItems] = useState<{ upc: string; product_number: number }[]>([]);
  const [products, setProducts] = useState<StoreProductWithDetails[]>([]);
  const [customerCards, setCustomerCards] = useState<CustomerCard[]>([]);
  const [employees, setEmployees] = useState<Employee[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<any>(null);

  useEffect(() => {
    fetchStoreProductsWithDetails().then(res => setProducts(res.data || []));
    getCustomerCards().then(setCustomerCards);
    fetchEmployees().then(res => setEmployees(res.data || []));
  }, []);

  const addItem = () => {
    setItems([...items, { upc: "", product_number: 1 }]);
  };

  const updateItem = (idx: number, field: "upc" | "product_number", value: any) => {
    setItems(items.map((item, i) => i === idx ? { ...item, [field]: value } : item));
  };

  const removeItem = (idx: number) => {
    setItems(items.filter((_, i) => i !== idx));
  };

  // Calculate price for an item
  const getItemPrice = (upc: string, quantity: number) => {
    const product = products.find(p => p.upc === upc);
    if (!product) return { price: 0, isPromo: false, total: 0 };
    let price = product.selling_price;
    let isPromo = product.promotional_product;
    if (isPromo) price = price * 0.8;
    return { price, isPromo, total: price * quantity };
  };

  const handleSubmit = async (e: any) => {
    e.preventDefault();
    setLoading(true);
    setError(null);
    setSuccess(null);
    try {
      const today = new Date();
      const yyyy = today.getFullYear();
      const mm = String(today.getMonth() + 1).padStart(2, '0');
      const dd = String(today.getDate()).padStart(2, '0');
      const print_date = `${yyyy}-${mm}-${dd}`;
      const check: CheckCreate = {
        employee_id: employeeId,
        card_number: cardNumber || undefined,
        print_date,
        items: items.map(i => {
          const { price } = getItemPrice(i.upc, i.product_number);
          return {
            upc: i.upc,
            product_number: Number(i.product_number),
            selling_price: price,
          };
        }),
      };
      const res = await createCheck(check);
      setSuccess(res.data);
      setItems([]);
    } catch (err: any) {
      setError(err?.response?.data?.error || "Failed to create check");
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
  const selectedCard = customerCards.find(card => card.card_number === cardNumber);
  const discountPercent = selectedCard ? selectedCard.percent : 0;
  const discountAmount = subtotal * (discountPercent / 100);
  const total = subtotal - discountAmount;

  return (
    <div className="max-w-2xl mx-auto p-6 bg-white rounded shadow">
      <h1 className="text-2xl font-bold mb-4">Create Check (Receipt)</h1>
      {error && <div className="bg-red-100 text-red-700 p-2 mb-2 rounded">{error}</div>}
      {success && (
        <div className="bg-green-100 text-green-700 p-2 mb-2 rounded">
          Check created! Receipt: {success.receipt_number}, Total: {success.total_sum}, VAT: {success.vat}
        </div>
      )}
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label>Cashier (Employee) *</label>
          <select
            value={employeeId}
            onChange={e => setEmployeeId(e.target.value)}
            className="w-full border p-2 rounded"
            required
          >
            <option value="">Select cashier...</option>
            {employees.map(emp => (
              <option key={emp.employee_id} value={emp.employee_id}>
                {emp.empl_surname} {emp.empl_name} {emp.empl_patronymic || ""} ({emp.employee_id})
              </option>
            ))}
          </select>
        </div>
        <div>
          <label>Customer Card (optional)</label>
          <select
            value={cardNumber || ""}
            onChange={e => setCardNumber(e.target.value || undefined)}
            className="w-full border p-2 rounded"
          >
            <option value="">No card</option>
            {customerCards.map(card => (
              <option key={card.card_number} value={card.card_number}>
                {card.cust_surname} {card.cust_name} {card.cust_patronymic || ""} ({card.card_number})
              </option>
            ))}
          </select>
        </div>
        <div>
          <label>Items</label>
          <button type="button" onClick={addItem} className="ml-2 px-2 py-1 bg-blue-500 text-white rounded">Add Item</button>
          {items.map((item, idx) => {
            const { price, isPromo, total } = getItemPrice(item.upc, item.product_number);
            const product = products.find(p => p.upc === item.upc);
            return (
              <div key={idx} className="flex gap-2 mt-2 items-center">
                <select value={item.upc} onChange={e => updateItem(idx, "upc", e.target.value)} className="border p-2 rounded">
                  <option value="">Select product</option>
                  {products.map(p => (
                    <option key={p.upc} value={p.upc}>{p.product_name} ({p.upc})</option>
                  ))}
                </select>
                <input type="number" min={1} value={item.product_number} onChange={e => updateItem(idx, "product_number", e.target.value)} className="w-20 border p-2 rounded" placeholder="Qty" />
                <span className="w-32">
                  {item.upc && product && (
                    <>
                      {isPromo ? (
                        <span className="text-orange-600 font-semibold">${price.toFixed(2)} (Promo -20%)</span>
                      ) : (
                        <span>${price.toFixed(2)}</span>
                      )}
                    </>
                  )}
                </span>
                <span className="w-24 text-right">{item.upc ? `Total: $${total.toFixed(2)}` : ""}</span>
                <button type="button" onClick={() => removeItem(idx)} className="text-red-500">Remove</button>
              </div>
            );
          })}
        </div>
        <div className="text-right font-bold text-lg">
          Subtotal: ${subtotal.toFixed(2)}<br />
          {discountPercent > 0 && (
            <span className="text-green-700">Discount ({discountPercent}%): -${discountAmount.toFixed(2)}</span>
          )}<br />
          Final Total: ${total.toFixed(2)}
        </div>
        <button type="submit" className="bg-green-600 text-white px-4 py-2 rounded" disabled={loading}>{loading ? "Creating..." : "Create Check"}</button>
      </form>
    </div>
  );
};

export default CreateCheck; 