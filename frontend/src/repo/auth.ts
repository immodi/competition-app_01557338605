import type { Token } from "@/interfaces/models/auth";
import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL;

async function signInUser(username: string, password: string): Promise<Token> {
    try {
        const response = await axios.post<Token>(`${API_URL}/auth/login`, {
            username,
            password,
        });

        return response.data;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message || `Login failed: ${error.message}`
        );
    }
}

async function registerUser(
    username: string,
    password: string
): Promise<Token> {
    try {
        const response = await axios.post<Token>(`${API_URL}/auth/register`, {
            username,
            password,
        });

        return response.data;
    } catch (error: any) {
        throw new Error(
            error.response?.data?.message || `Login failed: ${error.message}`
        );
    }
}

export { signInUser, registerUser };
