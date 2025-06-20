import { useState, useEffect } from 'react';
import { getCustomerCards, createCustomerCard, updateCustomerCard, deleteCustomerCard } from '../api/customer_cards';
import type { CustomerCard, CustomerCardCreate } from '../types/customer_card';
import { CustomerCard as CustomerCardComponent } from '../components/CustomerCard';

const CustomerCards = () => {
  const [customerCards, setCustomerCards] = useState<CustomerCard[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [isCreating, setIsCreating] = useState(false);
  const [newCard, setNewCard] = useState<CustomerCardCreate>({
    cust_surname: '',
    cust_name: '',
    cust_patronymic: '',
    phone_number: '',
    city: '',
    street: '',
    zip_code: '',
    percent: 0
  });
  const [search, setSearch] = useState("");
  const [sortField, setSortField] = useState<'cust_surname' | 'cust_name'>("cust_surname");
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>("asc");
  const [percentMin, setPercentMin] = useState<string>("");
  const [percentMax, setPercentMax] = useState<string>("");

  useEffect(() => {
    loadCustomerCards();
  }, []);

  const loadCustomerCards = async () => {
    try {
      setIsLoading(true);
      const cards = await getCustomerCards();
      setCustomerCards(cards || []);
      setError(null);
    } catch (err) {
      setError('Failed to load customer cards');
      console.error(err);
      setCustomerCards([]);
    } finally {
      setIsLoading(false);
    }
  };

  const handleCreate = async () => {
    try {
      await createCustomerCard(newCard);
      setIsCreating(false);
      setNewCard({
        cust_surname: '',
        cust_name: '',
        cust_patronymic: '',
        phone_number: '',
        city: '',
        street: '',
        zip_code: '',
        percent: 0
      });
      await loadCustomerCards();
      setError(null);
    } catch (err) {
      setError('Failed to create customer card');
      console.error(err);
    }
  };

  const handleUpdate = async (cardNumber: string, updates: Partial<CustomerCard>) => {
    try {
      await updateCustomerCard(cardNumber, updates);
      await loadCustomerCards();
      setError(null);
    } catch (err) {
      setError('Failed to update customer card');
      console.error(err);
    }
  };

  const handleDelete = async (cardNumber: string) => {
    try {
      await deleteCustomerCard(cardNumber);
      await loadCustomerCards();
      setError(null);
    } catch (err) {
      setError('Failed to delete customer card');
      console.error(err);
    }
  };

  const handleNewCardChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.name === 'percent' ? parseInt(e.target.value) : e.target.value;
    setNewCard({ ...newCard, [e.target.name]: value });
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold text-gray-900">Customer Cards</h1>
        <div className="flex gap-2 items-center">
          <input
            type="text"
            placeholder="Search by name or surname..."
            value={search}
            onChange={e => setSearch(e.target.value)}
            className="border rounded px-3 py-1"
            style={{ minWidth: 220 }}
          />
          <input
            type="number"
            placeholder="Min %"
            value={percentMin}
            onChange={e => setPercentMin(e.target.value)}
            className="border rounded px-2 py-1 w-20"
            min={0}
            max={100}
          />
          <input
            type="number"
            placeholder="Max %"
            value={percentMax}
            onChange={e => setPercentMax(e.target.value)}
            className="border rounded px-2 py-1 w-20"
            min={0}
            max={100}
          />
          <select value={sortField} onChange={e => setSortField(e.target.value as 'cust_surname' | 'cust_name')} className="border rounded px-2 py-1">
            <option value="cust_surname">Sort by Surname</option>
            <option value="cust_name">Sort by Name</option>
          </select>
          <select value={sortOrder} onChange={e => setSortOrder(e.target.value as 'asc' | 'desc')} className="border rounded px-2 py-1">
            <option value="asc">Asc</option>
            <option value="desc">Desc</option>
          </select>
        </div>
        <button
          className="bg-blue-500 text-white px-4 py-2 rounded"
          onClick={() => setIsCreating(true)}
        >
          Add New Card
        </button>
      </div>

      {error && (
        <div className="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
          {error}
        </div>
      )}

      {isCreating && (
        <div className="bg-white border border-gray-200 shadow-sm rounded-lg p-4 mb-4">
          <h2 className="text-lg font-semibold mb-4">New Customer Card</h2>
          <div className="grid grid-cols-1 gap-2 mb-4">
            <input name="cust_surname" value={newCard.cust_surname} onChange={handleNewCardChange} className="border p-1 rounded" placeholder="Surname" required />
            <input name="cust_name" value={newCard.cust_name} onChange={handleNewCardChange} className="border p-1 rounded" placeholder="Name" required />
            <input name="cust_patronymic" value={newCard.cust_patronymic} onChange={handleNewCardChange} className="border p-1 rounded" placeholder="Patronymic" />
            <input name="phone_number" value={newCard.phone_number} onChange={handleNewCardChange} className="border p-1 rounded" placeholder="Phone Number" pattern="\+380\d{9}" title="Format: +380XXXXXXXXX" required />
            <input name="city" value={newCard.city} onChange={handleNewCardChange} className="border p-1 rounded" placeholder="City" />
            <input name="street" value={newCard.street} onChange={handleNewCardChange} className="border p-1 rounded" placeholder="Street" />
            <input name="zip_code" value={newCard.zip_code} onChange={handleNewCardChange} className="border p-1 rounded" placeholder="Zip Code" />
            <input name="percent" type="number" value={newCard.percent} onChange={handleNewCardChange} className="border p-1 rounded" placeholder="Discount Percent" min="0" max="100" required />
          </div>
          <div className="flex gap-2 justify-end">
            <button className="bg-green-500 text-white px-3 py-1 rounded" onClick={handleCreate}>Create</button>
            <button className="bg-gray-300 text-gray-800 px-3 py-1 rounded" onClick={() => setIsCreating(false)}>Cancel</button>
          </div>
        </div>
      )}

      {isLoading ? (
        <div className="flex justify-center items-center py-8">
          <div className="text-gray-600">Loading customer cards...</div>
        </div>
      ) : customerCards.length === 0 ? (
        <div className="text-center py-8 text-gray-600">
          No customer cards found.
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
          {customerCards
            .filter(card => {
              const q = search.trim().toLowerCase();
              const min = percentMin === "" ? -Infinity : parseInt(percentMin);
              const max = percentMax === "" ? Infinity : parseInt(percentMax);
              const percentOk = card.percent >= min && card.percent <= max;
              return (
                percentOk &&
                (!q ||
                  card.cust_surname.toLowerCase().includes(q) ||
                  card.cust_name.toLowerCase().includes(q)
                )
              );
            })
            .sort((a, b) => {
              const fieldA = a[sortField].toLowerCase();
              const fieldB = b[sortField].toLowerCase();
              if (fieldA < fieldB) return sortOrder === 'asc' ? -1 : 1;
              if (fieldA > fieldB) return sortOrder === 'asc' ? 1 : -1;
              return 0;
            })
            .map((card) => (
              <CustomerCardComponent
                key={card.card_number}
                customerCard={card}
                onDelete={() => handleDelete(card.card_number)}
                onUpdate={(updates: Partial<CustomerCard>) => handleUpdate(card.card_number, updates)}
              />
            ))}
        </div>
      )}
    </div>
  );
};

export default CustomerCards; 