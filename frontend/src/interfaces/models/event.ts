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

interface EventsResponse {
    events: Event[];
    count: number;
}

interface AssignEventResponse {
    eventId: number;
}

interface CreateEventRequest {
    name: string;
    description: string;
    category: string;
    date: string;
    venue: string;
    price: number;
    image?: string;
    translations?: EventTranslation[];
}

export type {
    Event,
    EventsResponse,
    EventTranslation,
    AssignEventResponse,
    CreateEventRequest,
};
