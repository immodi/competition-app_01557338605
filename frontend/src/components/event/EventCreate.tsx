import { useNavigate } from "react-router";
import { useAppContext } from "@/contexts/AppContext";
import toast from "react-hot-toast";
import { createEvent } from "@/repo/events";
import useRequireAuth from "@/hooks/useRequireAuth";
import EventForm from "./EventForm";
import { IoArrowBack } from "react-icons/io5";
import type { CreateEventRequest } from "@/interfaces/models/event";

const EventCreatePage: React.FC = () => {
    const { authData } = useAppContext();
    const navigate = useNavigate();
    const isAuthed = useRequireAuth();

    if (!isAuthed)
        return <div className="p-10 text-center text-lg">Unauthorized</div>;

    async function handleCreateCallback(data: CreateEventRequest) {
        if (!authData.token) {
            toast.error("You must be logged in.");
            return;
        }

        // console.log(data);

        await createEvent(authData.token, data);
        toast.success("Event created!");
        navigate("/");
    }

    return (
        <main className="container mx-auto p-8 max-w-4xl">
            <div className="w-full flex items-center justify-between">
                <h1 className="text-3xl font-bold mb-6 text-gray-800 dark:text-gray-100">
                    Create Event
                </h1>
                <button
                    onClick={() => navigate(-1)}
                    className="flex cursor-pointer items-center gap-2 mb-6 px-4 py-2 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 dark:bg-gray-600 dark:text-gray-100 dark:hover:bg-gray-500 transition"
                >
                    <IoArrowBack />
                    Back
                </button>
            </div>
            <EventForm onSubmit={handleCreateCallback} submitLabel="Creatent" />
        </main>
    );
};

export default EventCreatePage;
