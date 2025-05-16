interface EventTranslation {
    language: string;
    name: string;
    description: string;
    venue: string;
}

interface Event {
    id: number;
    name: string;
    description: string;
    category: string;
    date: string;
    venue: string;
    price: number;
    image: string | null;
    translations: EventTranslation[];
}

interface TableProps {
    events: Event[];
}

export type { EventTranslation, Event, TableProps };
