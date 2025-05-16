import Login from "@/components/auth/Login";
import Register from "@/components/auth/Register";
import { createBrowserRouter } from "react-router";
import App from "../components/App";
import HomeRoute from "./HomeRoute";

const router = createBrowserRouter([
    {
        path: "/",
        Component: App,
        children: [
            { index: true, Component: HomeRoute },
            {
                path: "auth",
                children: [
                    { index: true, Component: Login },
                    { path: "login", Component: Login },
                    { path: "register", Component: Register },
                ],
            },
        ],
    },
]);

export default router;
