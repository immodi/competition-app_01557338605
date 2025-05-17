import { useLocation, useNavigate } from "react-router";
import { IoArrowBack } from "react-icons/io5";
import { TbConfetti } from "react-icons/tb";

const BookingSuccess: React.FC = () => {
    const navigate = useNavigate();

    const location = useLocation();
    const eventName = (location.state as { name?: string })?.name;

    return (
        <main className="container mx-auto p-8">
            <div className="w-full flex justify-between items-center">
                <div className="flex w-fit">
                    <TbConfetti size={40} className="relative bottom-1" />
                    <h1 className="text-3xl px-4 font-bold mb-6 text-gray-800 dark:text-gray-100">
                        Booking Confirmed!
                    </h1>
                </div>

                <button
                    onClick={() => navigate("/")}
                    className="flex cursor-pointer items-center gap-2 mb-6 px-4 py-2 bg-gray-200 text-gray-700 rounded hover:bg-gray-300 dark:bg-gray-600 dark:text-gray-100 dark:hover:bg-gray-500 transition"
                >
                    <IoArrowBack />
                    Home
                </button>
            </div>

            <div className="bg-white dark:bg-gray-800 rounded-lg shadow p-6 space-y-16 text-center">
                <h2 className="text-4xl font-semibold text-gray-800 dark:text-gray-100">
                    Congratulations!
                </h2>
                <p className="text-lg text-gray-700 dark:text-gray-300">
                    You've successfully booked{" "}
                    <span className="font-bold">
                        {eventName || "the event"}.
                    </span>
                </p>

                <button
                    onClick={() => navigate("/")}
                    className="mt-4 px-6 py-3 cursor-pointer bg-blue-600 text-white rounded-lg hover:bg-blue-700 dark:bg-blue-500 dark:hover:bg-blue-600 transition"
                >
                    Browse More Events
                </button>
            </div>
        </main>
    );
};

export default BookingSuccess;
