import type { Event } from "./models/event";

interface TableProps {
    events: Event[];
}

interface EventFormProps {
    initialData?: Partial<Event>;
    onSubmit: (data: any) => Promise<void>;
    submitLabel: string;
}

export type { TableProps, EventFormProps };
