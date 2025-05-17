import Login from "@/components/auth/Login";
import Register from "@/components/auth/Register";
import { createBrowserRouter } from "react-router";
import App from "../components/App";
import Home from "../components/home/Home";
import EventView from "@/components/event/EventView";
import EventCreate from "@/components/event/EventCreate";
import EventUpdate from "@/components/event/EventUpdate";
import BookingSuccess from "@/components/event/BookingSuccess";

const router = createBrowserRouter([
    {
        path: "/",
        Component: App,
        children: [
            { index: true, Component: Home },
            {
                path: ":category",
                Component: Home,
            },
            {
                path: "auth",
                children: [
                    { index: true, Component: Login },
                    { path: "login", Component: Login },
                    { path: "register", Component: Register },
                ],
            },
            {
                path: "event/:id",
                Component: EventView,
            },
            {
                path: "event/create",
                Component: EventCreate,
            },
            {
                path: "event/edit/:id",
                Component: EventUpdate,
            },
            {
                path: "booking-success",
                Component: BookingSuccess,
            },
        ],
    },
]);

export default router;
