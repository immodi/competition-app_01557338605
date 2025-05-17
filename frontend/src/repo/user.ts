import type { Event } from "@/interfaces/models/event";
import type { User } from "@/interfaces/models/user";
import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL;

function getAuthHeaders(token: string) {
    return {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
        "ngrok-skip-browser-warning": "true",
    };
}

async function getUserData(token: string): Promise<User> {
    try {
        const response = await axios.get<User>(`${API_URL}/users/data`, {
            headers: getAuthHeaders(token),
        });

        return response.data;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message || `Login failed: ${error.message}`
        );
    }
}

async function getUserEventIds(
    token: string,
    userId: number
): Promise<number[]> {
    try {
        const response = await axios.get<Event[]>(
            `${API_URL}/users/events/${userId}`,
            {
                headers: getAuthHeaders(token),
            }
        );

        const eventIds = response.data.map((event) => event.id);
        return eventIds;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message ||
                `Fetching user events failed: ${error.message}`
        );
    }
}

export { getUserData, getUserEventIds };
