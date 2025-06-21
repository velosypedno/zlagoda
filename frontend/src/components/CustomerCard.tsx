import { useState } from "react";
import type { CustomerCard as CustomerCardType } from "../types/customer_card";
import { useAuth } from "../contexts/AuthContext";

interface Props {
  customerCard: CustomerCardType;
  onDelete: () => void;
  onUpdate: (customerCard: Partial<CustomerCardType>) => void;
}

export const CustomerCard = ({ customerCard, onDelete, onUpdate }: Props) => {
  const { isManager } = useAuth();
  const [isEditing, setIsEditing] = useState(false);
  const [form, setForm] = useState<Partial<CustomerCardType>>({ ...customerCard });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.name === 'percent' ? parseInt(e.target.value) : e.target.value;
    setForm({ ...form, [e.target.name]: value });
  };

  const handleSave = () => {
    onUpdate(form);
    setIsEditing(false);
  };

  const handleCancel = () => {
    setForm({ ...customerCard });
    setIsEditing(false);
  };

  return (
    <div className="bg-white border border-gray-200 shadow-sm rounded-lg p-4">
      <div className="flex items-center justify-between mb-2">
        <div className="font-semibold text-lg text-gray-800">
          {customerCard.cust_surname} {customerCard.cust_name}
        </div>
        <span className="text-xs px-2 py-1 rounded bg-gray-100 text-gray-600 font-medium">
          {customerCard.percent}% Discount
        </span>
      </div>
      {isEditing ? (
        <div className="grid grid-cols-1 gap-2 mb-2">
          <input name="cust_surname" value={form.cust_surname || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Surname" />
          <input name="cust_name" value={form.cust_name || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Name" />
          <input name="cust_patronymic" value={form.cust_patronymic || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Patronymic" />
          <input name="phone_number" value={form.phone_number || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Phone Number" pattern="\+380\d{9}" title="Format: +380XXXXXXXXX" />
          <input name="city" value={form.city || ""} onChange={handleChange} className="border p-1 rounded" placeholder="City" />
          <input name="street" value={form.street || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Street" />
          <input name="zip_code" value={form.zip_code || ""} onChange={handleChange} className="border p-1 rounded" placeholder="Zip Code" />
          <input name="percent" type="number" value={form.percent || 0} onChange={handleChange} className="border p-1 rounded" placeholder="Discount Percent" min="0" max="100" />
        </div>
      ) : (
        <div className="text-sm text-gray-700 space-y-1 mb-2">
          <div><span className="font-medium">Card Number:</span> {customerCard.card_number}</div>
          <div><span className="font-medium">Phone:</span> {customerCard.phone_number}</div>
          <div><span className="font-medium">City:</span> {customerCard.city}</div>
          <div><span className="font-medium">Street:</span> {customerCard.street}</div>
          <div><span className="font-medium">Zip:</span> {customerCard.zip_code}</div>
        </div>
      )}
      
      {/* Only show action buttons for managers */}
      {isManager && (
        <div className="flex gap-2 mt-2 justify-end">
          {isEditing ? (
            <>
              <button className="bg-green-500 text-white px-3 py-1 rounded" onClick={handleSave}>Save</button>
              <button className="bg-gray-300 text-gray-800 px-3 py-1 rounded" onClick={handleCancel}>Cancel</button>
              <button className="bg-red-500 text-white px-3 py-1 rounded" onClick={onDelete}>Delete</button>
            </>
          ) : (
            <>
              <button className="bg-blue-500 text-white px-3 py-1 rounded" onClick={() => setIsEditing(true)}>Edit</button>
              <button className="bg-red-500 text-white px-3 py-1 rounded" onClick={onDelete}>Delete</button>
            </>
          )}
        </div>
      )}
    </div>
  );
};

export default CustomerCard; 