export interface SearchBarProps {
    placeholder?: string;
    query: string;

    setQuery: (query: string) => void;
    onSearch: (query: string) => void;
    onFocus: () => void;
}
