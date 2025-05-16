import Table from "@/components/home/Table";
import { useAppContext } from "@/contexts/AppContext";
import useRequireAuth from "@/hooks/useRequireAuth";
import type { Event } from "@/interfaces/event";
import { useEffect } from "react";

const events: Event[] = [
    {
        id: 1,
        name: "Test Event 1",
        description: "",
        category: "",
        date: "",
        venue: "",
        price: 100,
        image: null,
        translations: [],
    },
    {
        id: 2,
        name: "Test Event 2",
        description: "",
        category: "",
        date: "",
        venue: "",
        price: 80,
        image: null,
        translations: [],
    },
    {
        id: 3,
        name: "Test Event 3",
        description: "",
        category: "",
        date: "",
        venue: "",
        price: 50,
        image: null,
        translations: [],
    },
    {
        id: 4,
        name: "Test Event 4",
        description: "",
        category: "",
        date: "",
        venue: "",
        price: 120,
        image: null,
        translations: [],
    },
    {
        id: 5,
        name: "Test Event 5",
        description: "",
        category: "",
        date: "",
        venue: "",
        price: 95,
        image: null,
        translations: [],
    },
];
const HomeRoute: React.FC = () => {
    const { changeHeader } = useAppContext();
    const isAuthed = useRequireAuth();
    if (!isAuthed) {
        return null;
    }

    useEffect(() => {
        changeHeader("Events Dashboard");
    }, []);

    return (
        <main className="container mx-auto p-10">
            <Table events={events} />
        </main>
    );
};

export default HomeRoute;
