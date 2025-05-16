import type { User } from "./models/user";

export interface AuthInterface {
    token: string;
    isAuthed: boolean;
    userData: User;
    setToken: (token: string) => void;
    logout: () => void;
    refreshUserData: () => void;
}
