import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import { fetchCheckDetails, fetchCheckSalesWithDetails } from "../api/checks";
import { fetchEmployee } from "../api/employees";
import { getCustomerCard } from "../api/customer_cards";
import type { Check } from "../types/check";
import type { Employee } from "../types/employee";
import type { CustomerCard } from "../types/customer_card";
import type { SaleWithDetails } from "../types/sale";

const CheckDetails: React.FC = () => {
  const { receipt_number } = useParams<{ receipt_number: string }>();
  const [check, setCheck] = useState<Check | null>(null);
  const [cashier, setCashier] = useState<Employee | null>(null);
  const [customerCard, setCustomerCard] = useState<CustomerCard | null>(null);
  const [sales, setSales] = useState<SaleWithDetails[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!receipt_number) return;
    setLoading(true);
    setError(null);
    Promise.all([
      fetchCheckDetails(receipt_number),
      fetchCheckSalesWithDetails(receipt_number)
    ])
      .then(async ([checkRes, salesRes]) => {
        setCheck(checkRes.data);
        setSales(salesRes.data);
        // Fetch cashier
        const cashierRes = await fetchEmployee(checkRes.data.employee_id);
        setCashier(cashierRes.data);
        // Fetch customer card if present
        if (checkRes.data.card_number) {
          try {
            const cardRes = await getCustomerCard(checkRes.data.card_number);
            setCustomerCard(cardRes);
          } catch {
            setCustomerCard(null);
          }
        } else {
          setCustomerCard(null);
        }
      })
      .catch((err) => {
        setError(err?.response?.data?.error || err.message || "Failed to load check details");
      })
      .finally(() => setLoading(false));
  }, [receipt_number]);

  if (loading) return <div className="p-8">Loading...</div>;
  if (error) return <div className="p-8 text-red-600">{error}</div>;
  if (!check) return <div className="p-8">Check not found.</div>;

  return (
    <div className="max-w-lg mx-auto p-6 bg-white rounded shadow mt-6">
      <div className="mb-4 flex justify-between items-center">
        <h2 className="text-xl font-bold">Receipt #{check.receipt_number}</h2>
        <Link to="/checks" className="text-blue-600 hover:underline">Back to list</Link>
      </div>
      <div className="mb-2 text-sm text-gray-600">Date: {check.print_date}</div>
      <div className="mb-2 text-sm text-gray-600">
        <span className="font-semibold">Cashier:</span>{" "}
        {cashier ? (
          <span>{cashier.empl_surname} {cashier.empl_name} {cashier.empl_patronymic || ""} ({cashier.empl_role})</span>
        ) : (
          <span className="text-gray-400">Unknown</span>
        )}
      </div>
      {customerCard && (
        <div className="mb-2 text-sm text-gray-600">
          <span className="font-semibold">Customer Card:</span> {customerCard.cust_surname} {customerCard.cust_name} ({customerCard.card_number}) — <span className="font-semibold">{customerCard.percent}% discount</span>
        </div>
      )}
      <div className="border-b my-4" />
      <div className="mb-2 font-semibold">Items:</div>
      <table className="w-full text-sm mb-4">
        <thead>
          <tr className="border-b">
            <th className="text-left py-1">Product</th>
            <th className="text-right py-1">Qty</th>
            <th className="text-right py-1">Price</th>
            <th className="text-right py-1">Total</th>
          </tr>
        </thead>
        <tbody>
          {sales.map((item) => (
            <tr key={item.upc}>
              <td className="py-1">
                <div className="font-medium">{item.product_name}</div>
                <div className="text-xs text-gray-500">{item.category_name}</div>
                <div className="text-xs text-gray-400">{item.characteristics}</div>
              </td>
              <td className="text-right py-1">{item.product_number}</td>
              <td className="text-right py-1">{item.selling_price.toFixed(2)}</td>
              <td className="text-right py-1">{item.total_price.toFixed(2)}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <div className="border-b my-4" />
      <div className="flex justify-between text-sm mb-1">
        <span>Subtotal:</span>
        <span>{check.sum_total.toFixed(2)}</span>
      </div>
      {customerCard && customerCard.percent > 0 && (
        <div className="flex justify-between text-sm mb-1">
          <span>Discount ({customerCard.percent}%):</span>
          <span>-{(check.sum_total * customerCard.percent / 100).toFixed(2)}</span>
        </div>
      )}
      <div className="flex justify-between text-sm mb-1">
        <span>VAT:</span>
        <span>{check.vat.toFixed(2)}</span>
      </div>
      <div className="flex justify-between text-base font-bold mt-2">
        <span>Total:</span>
        <span>
          {customerCard && customerCard.percent > 0
            ? (check.sum_total - (check.sum_total * customerCard.percent / 100)).toFixed(2)
            : check.sum_total.toFixed(2)}
        </span>
      </div>
    </div>
  );
};

export default CheckDetails; 