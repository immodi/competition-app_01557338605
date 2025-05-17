import type { SearchBarProps } from "@/interfaces/searchBar";
import { useEffect, useState } from "react";
import { FaTimes } from "react-icons/fa";
import { useNavigate } from "react-router";

const SearchBar: React.FC<SearchBarProps> = ({
    placeholder = "Search events...",
    query,

    onSearch,
    onFocus,
    setQuery,
}) => {
    const navigate = useNavigate();
    const [isDisabled, setIsDisabled] = useState(false);

    useEffect(() => {
        //?category=Test
        if (query.slice(0, 10) === "?category=") {
            navigate(`/${query.slice(10)}`);
            setIsDisabled(true);
        }

        if (query.charAt(0) !== "?") {
            onSearch(query);
        }
    }, [query]);

    return (
        <div className="flex items-center gap-2 mb-4">
            <input
                type="text"
                value={query}
                onFocus={onFocus}
                onChange={(e) => setQuery(e.target.value)}
                placeholder={placeholder}
                disabled={isDisabled}
                className="px-4 py-2 rounded border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-800 text-gray-800 dark:text-gray-100 focus:outline-none focus:ring-2 focus:ring-blue-500 w-full"
            />

            <button
                onClick={() => {
                    navigate("/");
                    setIsDisabled(false);
                    setQuery("");
                }}
                className="px-3 py-2 cursor-pointer rounded bg-gray-300 text-gray-800 hover:bg-gray-400 dark:bg-gray-700 dark:text-gray-100 dark:hover:bg-gray-600 transition"
            >
                <FaTimes />
            </button>
        </div>
    );
};

export default SearchBar;
