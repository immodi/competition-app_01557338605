import { useAppContext } from "@/contexts/AppContext";
import type { TableProps } from "@/interfaces/tableData";
import { assignUserToEvent } from "@/repo/events";
import { getUserEventIds } from "@/repo/user";
import { useEffect, useState } from "react";
import type React from "react";
import toast from "react-hot-toast";
import { FaPencil, FaTicket } from "react-icons/fa6";
import { IoEyeSharp } from "react-icons/io5";
import { useNavigate } from "react-router";

const Table: React.FC<TableProps> = ({ events }) => {
    const { authData } = useAppContext();
    const [userEventIds, setUserEventIds] = useState<number[]>([]);
    const navigate = useNavigate();

    useEffect(() => {
        if (authData.token && authData.userData.userId !== 0) {
            getUserEventIds(authData.token, authData.userData.userId)
                .then(setUserEventIds)
                .catch((error) => {
                    console.error(error);
                    toast.error(error.message);
                });
        }
    }, [authData.token, authData.userData.userId, authData.userData.tickets]);

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
                        ].map((heading) => (
                            <th
                                key={heading}
                                className="px-4 py-3 text-left text-gray-700 dark:text-gray-200 font-semibold border-b dark:border-gray-600"
                            >
                                {heading}
                            </th>
                        ))}
                        <th className="px-4 py-3 text-center text-gray-700 dark:text-gray-200 font-semibold border-b dark:border-gray-600">
                            Actions
                        </th>
                    </tr>
                </thead>
                <tbody>
                    {events.map((event) => (
                        <tr
                            key={event.id}
                            className="hover:bg-gray-50 border-b dark:border-gray-600 dark:hover:bg-gray-700"
                        >
                            <td className="px-4 py-3 dark:text-gray-100">
                                {event.id}
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                {event.name}
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                {event.description}
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                {event.category}
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                {new Date(event.date).toLocaleString()}
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                {event.venue}
                            </td>
                            <td className="px-4 py-3 dark:text-gray-100">
                                ${event.price}
                            </td>
                            <td className="px-4 py-3 items-center ">
                                <div className="flex justify-center gap-2">
                                    {!userEventIds.includes(event.id) ? (
                                        <button
                                            onClick={() => {
                                                assignUserToEvent(
                                                    authData.token,
                                                    event.id,
                                                    authData.userData.userId
                                                )
                                                    .then(() => {
                                                        authData.refreshUserData();
                                                        toast.success(
                                                            "event booked Successfully"
                                                        );
                                                    })
                                                    .catch((error) => {
                                                        console.error(error);
                                                        toast.error(
                                                            error.message
                                                        );
                                                    });
                                            }}
                                            title="book button"
                                            className="px-3 py-1 bg-blue-500 cursor-pointer text-white rounded hover:bg-blue-600 transition"
                                        >
                                            <FaTicket />
                                        </button>
                                    ) : (
                                        <button
                                            disabled
                                            title="booked"
                                            className="px-3 py-1 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 dark:bg-gray-600 dark:text-gray-100 dark:hover:bg-gray-500 transition"
                                        >
                                            <FaTicket />
                                        </button>
                                    )}
                                    <button
                                        onClick={() =>
                                            navigate(`/event/${event.id}`)
                                        }
                                        title="view details"
                                        className="px-3 py-1 bg-blue-600 text-white cursor-pointer rounded hover:bg-blue-700 dark:bg-blue-500 dark:hover:bg-blue-600 transition"
                                    >
                                        <IoEyeSharp />
                                    </button>
                                    {authData.userData.role === "admin" && (
                                        <button
                                            onClick={() =>
                                                navigate(
                                                    `/event/edit/${event.id}`
                                                )
                                            }
                                            title="edit event"
                                            className="px-3 py-1 bg-green-600 text-white cursor-pointer rounded hover:bg-green-700 dark:bg-green-500 dark:hover:bg-green-600 transition"
                                        >
                                            <FaPencil />
                                        </button>
                                    )}
                                </div>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default Table;
