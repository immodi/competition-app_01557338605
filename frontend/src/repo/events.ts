import type {
    AssignEventResponse,
    CreateEventRequest,
    Event,
    EventsResponse,
} from "@/interfaces/models/event";
import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL;

async function getEvents(
    token: string,
    page: number = 1,
    limit: number = 10
): Promise<EventsResponse> {
    try {
        const response = await axios.get<EventsResponse>(`${API_URL}/events`, {
            headers: {
                Authorization: `Bearer ${token}`,
                "Content-Type": "application/json",
            },
            params: {
                page,
                limit,
            },
        });

        return response.data;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message ||
                `Fetching events failed: ${error.message}`
        );
    }
}

async function assignUserToEvent(
    token: string,
    eventId: number,
    userId: number
): Promise<AssignEventResponse> {
    try {
        const response = await axios.post<AssignEventResponse>(
            `${API_URL}/events/assign/${eventId}`,
            { userId },
            {
                headers: {
                    Authorization: `Bearer ${token}`,
                    "Content-Type": "application/json",
                },
            }
        );

        return response.data;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message ||
                `Assigning user to event failed: ${error.message}`
        );
    }
}

async function getEventById(token: string, eventId: number): Promise<Event> {
    try {
        const response = await axios.get<Event>(
            `${API_URL}/events/${eventId}`,
            {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            }
        );

        return response.data;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message ||
                `Fetching event failed: ${error.message}`
        );
    }
}

async function createEvent(
    token: string,
    eventData: CreateEventRequest
): Promise<Event> {
    try {
        const response = await axios.post<Event>(
            `${API_URL}/events`,
            eventData,
            {
                headers: {
                    Authorization: `Bearer ${token}`,
                    "Content-Type": "application/json",
                },
            }
        );

        return response.data;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message ||
                `Creating event failed: ${error.message}`
        );
    }
}

async function updateEvent(
    token: string,
    eventId: number,
    eventData: CreateEventRequest
): Promise<Event> {
    try {
        const response = await axios.put<Event>(
            `${API_URL}/events/${eventId}`,
            eventData,
            {
                headers: {
                    Authorization: `Bearer ${token}`,
                    "Content-Type": "application/json",
                },
            }
        );

        return response.data;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message ||
                `Updating event failed: ${error.message}`
        );
    }
}

async function getEventsByCategory(
    token: string,
    category: string,
    page: number = 1,
    limit: number = 10
): Promise<EventsResponse> {
    try {
        const response = await axios.get<EventsResponse>(
            `${API_URL}/events/category/${category}`,
            {
                headers: {
                    Authorization: `Bearer ${token}`,
                    "Content-Type": "application/json",
                },
                params: {
                    page,
                    limit,
                },
            }
        );

        return response.data;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message ||
                `Fetching events failed: ${error.message}`
        );
    }
}

async function searchEvents(
    token: string,
    query: string,
    page: number = 1,
    limit: number = 10
): Promise<EventsResponse> {
    try {
        const response = await axios.get<EventsResponse>(
            `${API_URL}/events/search/${query}`,
            {
                headers: {
                    Authorization: `Bearer ${token}`,
                    "Content-Type": "application/json",
                },
                params: {
                    page,
                    limit,
                },
            }
        );

        return response.data;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message ||
                `Fetching events failed: ${error.message}`
        );
    }
}

async function deleteEvent(token: string, eventId: number): Promise<void> {
    try {
        await axios.delete(`${API_URL}/events/${eventId}`, {
            headers: {
                Authorization: `Bearer ${token}`,
                "Content-Type": "application/json",
            },
        });
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message ||
                `Deleting event failed: ${error.message}`
        );
    }
}

export {
    getEvents,
    assignUserToEvent,
    getEventById,
    createEvent,
    getEventsByCategory,
    searchEvents,
    updateEvent,
    deleteEvent,
};
