import type { TableProps } from "@/interfaces/event";
import type React from "react";
import { FaTicket } from "react-icons/fa6";
import { IoEyeSharp } from "react-icons/io5";

const Table: React.FC<TableProps> = ({ events }) => {
    return (
        <div className="overflow-x-auto">
            <table className="min-w-full border border-gray-200 bg-white dark:bg-gray-800 dark:border-gray-700 rounded-lg">
                <thead className="bg-gray-100 dark:bg-gray-700">
                    <tr>
                        {[
                            "Event Id",
                            "Event Name",
                            "Description",
                            "Category",
                            "Date",
                            "Venue",
                            "Price",
                            "Actions",
                        ].map((heading) => (
                            <th
                                key={heading}
                                className="px-4 py-3 text-left text-gray-700 dark:text-gray-200 font-semibold border-b dark:border-gray-600"
                            >
                                {heading}
                            </th>
                        ))}
                    </tr>
                </thead>
                <tbody>
                    {events.map((event) => (
                        <tr
                            key={event.id}
                            className="hover:bg-gray-50 dark:hover:bg-gray-700"
                        >
                            <td className="px-4 py-3 border-b dark:border-gray-600 dark:text-gray-100">
                                {event.id}
                            </td>
                            <td className="px-4 py-3 border-b dark:border-gray-600 dark:text-gray-100">
                                {event.name}
                            </td>
                            <td className="px-4 py-3 border-b dark:border-gray-600 dark:text-gray-100">
                                {event.description}
                            </td>
                            <td className="px-4 py-3 border-b dark:border-gray-600 dark:text-gray-100">
                                {event.category}
                            </td>
                            <td className="px-4 py-3 border-b dark:border-gray-600 dark:text-gray-100">
                                {new Date(event.date).toLocaleString()}
                            </td>
                            <td className="px-4 py-3 border-b dark:border-gray-600 dark:text-gray-100">
                                {event.venue}
                            </td>
                            <td className="px-4 py-3 border-b dark:border-gray-600 dark:text-gray-100">
                                ${event.price}
                            </td>
                            <td className="px-4 py-3 border-b flex gap-2 dark:border-gray-600">
                                <button
                                    onClick={() =>
                                        alert(`Booking event: ${event.name}`)
                                    }
                                    title="book button"
                                    className="px-3 py-1 bg-blue-500 cursor-pointer text-white rounded hover:bg-blue-600 transition"
                                >
                                    <FaTicket />
                                </button>
                                <button
                                    onClick={() =>
                                        alert(
                                            `Viewing details for: ${event.name}`
                                        )
                                    }
                                    title="view details"
                                    className="px-3 py-1 bg-gray-200 text-gray-700 cursor-pointer rounded hover:bg-gray-300 dark:bg-gray-600 dark:text-gray-100 dark:hover:bg-gray-500 transition"
                                >
                                    <IoEyeSharp />
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default Table;
