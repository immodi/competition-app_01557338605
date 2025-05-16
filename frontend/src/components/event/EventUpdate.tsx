import { useParams, useNavigate } from "react-router";
import { useEffect, useState } from "react";
import { getEventById, updateEvent } from "@/repo/events";
import { useAppContext } from "@/contexts/AppContext";
import toast from "react-hot-toast";
import useRequireAuth from "@/hooks/useRequireAuth";
import EventForm from "./EventForm";
import type { Event } from "@/interfaces/models/event";
import { IoArrowBack } from "react-icons/io5";

const EventUpdate: React.FC = () => {
    const { id } = useParams<{ id: string }>();
    const { authData } = useAppContext();
    const navigate = useNavigate();
    const isAuthed = useRequireAuth();
    const [eventData, setEventData] = useState<Event>({} as Event);

    useEffect(() => {
        if (!authData.token) return;
        getEventById(authData.token, Number(id))
            .then(setEventData)
            .catch((error) => {
                console.error(error);
                toast.error(error.message);
            });
    }, [id, authData.token]);

    if (!isAuthed)
        return <div className="p-10 text-center text-lg">Unauthorized</div>;
    if (!eventData)
        return <div className="p-10 text-center text-lg">Loading event...</div>;

    const handleUpdate = async (data: any) => {
        if (!authData.token) {
            toast.error("You must be logged in.");
            return;
        }
        await updateEvent(authData.token, Number(id), data);
        toast.success("Event updated!");
        navigate("/");
    };

    return (
        <main className="container mx-auto p-8 max-w-4xl">
            <div className="w-full flex items-center justify-between">
                <h1 className="text-3xl font-bold mb-6 text-gray-800 dark:text-gray-100">
                    Update Event
                </h1>
                <button
                    onClick={() => navigate(-1)}
                    className="flex cursor-pointer items-center gap-2 mb-6 px-4 py-2 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 dark:bg-gray-600 dark:text-gray-100 dark:hover:bg-gray-500 transition"
                >
                    <IoArrowBack />
                    Back
                </button>
            </div>
            <EventForm
                initialData={eventData}
                onSubmit={handleUpdate}
                submitLabel="Update Event"
            />
        </main>
    );
};

export default EventUpdate;
