export interface User {
    userId: number;
    username: string;
    createdAt: Date;
    tickets: number;
    role: "admin" | "user";
}
