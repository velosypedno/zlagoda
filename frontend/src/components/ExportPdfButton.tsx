import React, { useState } from "react";

interface ExportPdfButtonProps {
  entityType: string;
  entityId?: string | number;
  apiEndpoint: string;
  title?: string;
  columns: Array<{
    key: string;
    label: string;
    width?: string | number;
  }>;
  customStyles?: Record<string, unknown>;
  className?: string;
  children?: React.ReactNode;
}

const ExportPdfButton: React.FC<ExportPdfButtonProps> = ({
  entityType,
  entityId,
  apiEndpoint,
  title,
  columns,
  className = "bg-green-500 text-white px-4 py-2 rounded hover:bg-green-600 transition",
  children,
}) => {
  const [isExporting, setIsExporting] = useState(false);

  const fetchData = async (): Promise<unknown> => {
    const token = localStorage.getItem("token");
    const baseUrl =
      import.meta.env.VITE_FRONTEND_BASE_API_URL || "http://localhost:8080";

    const response = await fetch(`${baseUrl}${apiEndpoint}`, {
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error(`Failed to fetch ${entityType} data`);
    }

    return response.json();
  };

  const formatValue = (value: unknown): string => {
    if (value === null || value === undefined) {
      return "";
    }
    if (typeof value === "boolean") {
      return value ? "Yes" : "No";
    }
    if (typeof value === "number") {
      return value.toString();
    }
    if (typeof value === "object") {
      return JSON.stringify(value);
    }
    return String(value);
  };

  const getNestedValue = (
    obj: Record<string, unknown>,
    path: string,
  ): unknown => {
    const keys = path.split(".");
    let current: unknown = obj;

    for (const key of keys) {
      if (
        current &&
        typeof current === "object" &&
        key in (current as Record<string, unknown>)
      ) {
        current = (current as Record<string, unknown>)[key];
      } else {
        return "";
      }
    }

    return current;
  };

  const createTableFromData = (data: unknown): string[][] => {
    const items = Array.isArray(data) ? data : [data];
    const headers = columns.map((col) => col.label);

    const rows = items.map((item) =>
      columns.map((col) => {
        const value = getNestedValue(item as Record<string, unknown>, col.key);
        return formatValue(value);
      }),
    );

    return [headers, ...rows];
  };

  const generatePrintWindow = async (data: unknown) => {
    const tableData = createTableFromData(data);
    const currentDate = new Date().toLocaleDateString();
    const currentTime = new Date().toLocaleTimeString();

    const htmlContent = `
      <!DOCTYPE html>
      <html>
      <head>
        <meta charset="utf-8">
        <title>${title || `${entityType} Export`}</title>
        <style>
          body {
            font-family: Arial, sans-serif;
            margin: 20px;
            font-size: 12px;
          }
          .header {
            display: flex;
            justify-content: space-between;
            margin-bottom: 20px;
            padding-bottom: 10px;
            border-bottom: 2px solid #3b82f6;
          }
          .title {
            font-size: 24px;
            font-weight: bold;
            text-align: center;
            margin: 20px 0;
            color: #1f2937;
          }
          .subtitle {
            font-size: 14px;
            text-align: center;
            margin-bottom: 20px;
            color: #6b7280;
          }
          table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 20px;
          }
          th, td {
            border: 1px solid #e5e7eb;
            padding: 8px;
            text-align: left;
            font-size: 11px;
          }
          th {
            background-color: #f3f4f6;
            font-weight: bold;
            color: #374151;
          }
          tr:nth-child(even) {
            background-color: #f9fafb;
          }
          .buttons {
            margin-top: 20px;
            text-align: center;
            page-break-inside: avoid;
          }
          .btn {
            padding: 10px 20px;
            margin: 0 5px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            font-size: 14px;
          }
          .btn-print {
            background: #3b82f6;
            color: white;
          }
          .btn-close {
            background: #6b7280;
            color: white;
          }
          @media print {
            .buttons { display: none; }
            body { margin: 0; }
          }
        </style>
      </head>
      <body>
        <div class="header">
          <div>🏪 Zlagoda Store Management</div>
          <div>Generated: ${currentDate} ${currentTime}</div>
        </div>

        <div class="title">${title || `${entityType} Export`}</div>

        ${entityId ? `<div class="subtitle">${entityType} ID: ${entityId}</div>` : ""}

        <table>
          <thead>
            <tr>
              ${tableData[0].map((header) => `<th>${header}</th>`).join("")}
            </tr>
          </thead>
          <tbody>
            ${tableData
              .slice(1)
              .map(
                (row) =>
                  `<tr>${row.map((cell) => `<td>${cell}</td>`).join("")}</tr>`,
              )
              .join("")}
          </tbody>
        </table>

        <div class="buttons">
          <button class="btn btn-print" onclick="window.print()">
            🖨️ Print / Save as PDF
          </button>
          <button class="btn btn-close" onclick="window.close()">
            ✖️ Close
          </button>
        </div>

        <script>
          // Auto-focus the window
          window.focus();
        </script>
      </body>
      </html>
    `;

    const printWindow = window.open("", "_blank", "width=800,height=600");
    if (printWindow) {
      printWindow.document.write(htmlContent);
      printWindow.document.close();
    } else {
      throw new Error(
        "Could not open print window. Please check popup settings.",
      );
    }
  };

  const handleExport = async () => {
    try {
      setIsExporting(true);
      const data = await fetchData();
      await generatePrintWindow(data);
    } catch (error) {
      console.error("Export error:", error);
      const errorMessage =
        error instanceof Error ? error.message : "Unknown error";
      alert(`Failed to export ${entityType}: ${errorMessage}`);
    } finally {
      setIsExporting(false);
    }
  };

  return (
    <button
      onClick={handleExport}
      disabled={isExporting}
      className={className}
      title={`Export ${entityType} to PDF`}
    >
      {isExporting ? (
        <>
          <span className="animate-spin inline-block w-4 h-4 border-2 border-white border-t-transparent rounded-full mr-2"></span>
          Exporting...
        </>
      ) : (
        children || <>📄 Export PDF</>
      )}
    </button>
  );
};

export default ExportPdfButton;
