import React, { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import {
  fetchReceiptDetails,
  fetchReceiptSalesWithDetails,
} from "../api/receipts";
import { fetchEmployee } from "../api/employees";
import { getCustomerCard } from "../api/customer_cards";
import type { ReceiptRetrieve } from "../types/receipt";
import type { Employee } from "../types/employee";
import type { CustomerCard } from "../types/customer_card";
import type { SaleWithDetails } from "../types/sale";

const ReceiptDetails: React.FC = () => {
  const { receipt_number } = useParams<{ receipt_number: string }>();
  const [receipt, setReceipt] = useState<ReceiptRetrieve | null>(null);
  const [cashier, setCashier] = useState<Employee | null>(null);
  const [customerCard, setCustomerCard] = useState<CustomerCard | null>(null);
  const [sales, setSales] = useState<SaleWithDetails[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const loadReceiptData = async () => {
      if (!receipt_number) {
        setError("No receipt number provided");
        setLoading(false);
        return;
      }

      try {
        setLoading(true);
        setError(null);

        // Fetch receipt details
        const receiptRes = await fetchReceiptDetails(receipt_number);
        setReceipt(receiptRes.data);

        // Fetch sales details
        try {
          const salesRes = await fetchReceiptSalesWithDetails(receipt_number);
          setSales(salesRes.data);
        } catch (err) {
          console.error("Failed to fetch sales:", err);
          setSales([]);
        }

        // Fetch cashier info
        if (receiptRes.data.employee_id) {
          try {
            const cashierRes = await fetchEmployee(receiptRes.data.employee_id);
            setCashier(cashierRes.data);
          } catch (err) {
            console.error("Failed to fetch cashier:", err);
            setCashier(null);
          }
        }

        // Fetch customer card if present
        if (receiptRes.data.card_number) {
          try {
            const cardRes = await getCustomerCard(receiptRes.data.card_number);
            setCustomerCard(cardRes);
          } catch (err) {
            console.error("Failed to fetch customer card:", err);
            setCustomerCard(null);
          }
        }
      } catch (err: any) {
        console.error("Error loading receipt:", err);
        setError(
          err?.response?.data?.error ||
            err.message ||
            "Failed to load receipt details",
        );
      } finally {
        setLoading(false);
      }
    };

    loadReceiptData();
  }, [receipt_number]);

  if (loading) {
    return (
      <div className="max-w-lg mx-auto p-6 bg-white rounded shadow mt-6">
        <div className="text-center">Loading receipt details...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="max-w-lg mx-auto p-6 bg-white rounded shadow mt-6">
        <div className="mb-4 flex justify-between items-center">
          <h2 className="text-xl font-bold text-red-600">Error</h2>
          <Link to="/receipts" className="text-blue-600 hover:underline">
            Back to list
          </Link>
        </div>
        <div className="text-red-600">{error}</div>
      </div>
    );
  }

  if (!receipt) {
    return (
      <div className="max-w-lg mx-auto p-6 bg-white rounded shadow mt-6">
        <div className="mb-4 flex justify-between items-center">
          <h2 className="text-xl font-bold">Receipt Not Found</h2>
          <Link to="/receipts" className="text-blue-600 hover:underline">
            Back to list
          </Link>
        </div>
        <div>Receipt not found.</div>
      </div>
    );
  }

  // Format timestamp to show date and time
  const formatTimestamp = (dateString: string) => {
    try {
      const date = new Date(dateString);
      return date.toLocaleString("en-US", {
        year: "numeric",
        month: "2-digit",
        day: "2-digit",
        hour: "2-digit",
        minute: "2-digit",
        second: "2-digit",
        hour12: false,
      });
    } catch {
      return dateString; // fallback to original string if parsing fails
    }
  };

  return (
    <div className="max-w-lg mx-auto p-6 bg-white rounded shadow mt-6">
      <div className="mb-4 flex justify-between items-center">
        <h2 className="text-xl font-bold">Receipt #{receipt.receipt_number}</h2>
        <Link to="/receipts" className="text-blue-600 hover:underline">
          Back to list
        </Link>
      </div>

      <div className="mb-2 text-sm text-gray-600">
        Timestamp: {formatTimestamp(receipt.print_date)}
      </div>

      <div className="mb-2 text-sm text-gray-600">
        <span className="font-semibold">Cashier:</span>{" "}
        {cashier ? (
          <span>
            {cashier.empl_surname || ""} {cashier.empl_name || ""}{" "}
            {cashier.empl_patronymic || ""} ({cashier.empl_role || ""})
          </span>
        ) : (
          <span className="text-gray-400">
            Employee ID: {receipt.employee_id}
          </span>
        )}
      </div>

      {customerCard && (
        <div className="mb-2 text-sm text-gray-600">
          <span className="font-semibold">Customer Card:</span>{" "}
          {customerCard.cust_surname} {customerCard.cust_name} (
          {customerCard.card_number}) —{" "}
          <span className="font-semibold">
            {customerCard.percent}% discount
          </span>
        </div>
      )}

      <div className="border-b my-4" />

      <div className="mb-2 font-semibold">Items:</div>
      {sales && sales.length > 0 ? (
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
                  <div className="text-xs text-gray-500">
                    {item.category_name}
                  </div>
                  <div className="text-xs text-gray-400">
                    {item.characteristics}
                  </div>
                </td>
                <td className="text-right py-1">{item.product_number}</td>
                <td className="text-right py-1">
                  ${item.selling_price.toFixed(2)}
                </td>
                <td className="text-right py-1">
                  ${item.total_price.toFixed(2)}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <div className="text-gray-500 mb-4">
          No items found for this receipt.
        </div>
      )}

      <div className="border-b my-4" />

      <div className="flex justify-between text-sm mb-1">
        <span>Subtotal:</span>
        <span>${receipt.sum_total?.toFixed(2) || "0.00"}</span>
      </div>

      {customerCard && customerCard.percent > 0 && (
        <div className="flex justify-between text-sm mb-1">
          <span>Discount ({customerCard.percent}%):</span>
          <span>
            -$
            {(((receipt.sum_total || 0) * customerCard.percent) / 100).toFixed(
              2,
            )}
          </span>
        </div>
      )}

      <div className="flex justify-between text-sm mb-1">
        <span>VAT:</span>
        <span>${receipt.vat?.toFixed(2) || "0.00"}</span>
      </div>

      <div className="flex justify-between text-base font-bold mt-2">
        <span>Total:</span>
        <span>
          $
          {customerCard && customerCard.percent > 0
            ? (
                (receipt.sum_total || 0) -
                ((receipt.sum_total || 0) * customerCard.percent) / 100
              ).toFixed(2)
            : (receipt.sum_total || 0).toFixed(2)}
        </span>
      </div>
    </div>
  );
};

export default ReceiptDetails;
