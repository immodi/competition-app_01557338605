import type { Event } from "./models/event";

interface TableProps {
    events: Event[];
    filterCategory: (category: string) => void;
    refreshEvents: () => void;
}

interface EventFormProps {
    initialData?: Partial<Event>;
    onSubmit: (data: any) => Promise<void>;
    submitLabel: string;
}

export type { TableProps, EventFormProps };
