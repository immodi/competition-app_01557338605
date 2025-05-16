import toBase64 from "@/helpers/toBase64";
import type { EventTranslation } from "@/interfaces/models/event";
import type { EventFormProps } from "@/interfaces/tableData";
import { useEffect, useState } from "react";
import toast from "react-hot-toast";

const EventForm: React.FC<EventFormProps> = ({
    initialData = {},
    onSubmit,
    submitLabel,
}) => {
    const [data, setData] = useState(initialData);
    useEffect(() => {
        if (initialData && Object.keys(initialData).length > 0) {
            setData(initialData);
        }
    }, [initialData]);
    const handleAddTranslation = () => {
        // setTranslations([
        //     ...translations,
        //     { language: "", name: "", description: "", venue: "" },
        // ]);
        setData({
            ...data,
            translations: [
                ...(data.translations || []),
                {
                    language: "",
                    name: "",
                    description: "",
                    venue: "",
                },
            ],
        });
    };

    const handleRemoveTranslation = (index: number) => {
        // setTranslations(translations.filter((_, i) => i !== index));
        setData({
            ...data,
            translations: data.translations?.filter((_, i) => i !== index),
        });
    };

    const handleTranslationChange = (
        index: number,
        field: keyof EventTranslation,
        value: string
    ) => {
        // const updated = [...translations];
        // updated[index][field] = value;
        // setTranslations(updated);

        setData({
            ...data,
            translations: data.translations?.map((t, i) =>
                i === index ? { ...t, [field]: value } : t
            ),
        });
    };

    const handleImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files[0]) {
            // setImageFile(e.target.files[0]);
            toBase64(e.target.files[0])
                .then((base64) => {
                    setData({
                        ...data,
                        image: base64,
                    });
                })
                .catch((error) => {
                    toast.error(error.message || "Something went wrong.");
                });
        }
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            await onSubmit(data);
        } catch (err: any) {
            toast.error(err.message || "Something went wrong.");
        }
    };

    return (
        <form
            onSubmit={handleSubmit}
            className="space-y-6 bg-white dark:bg-gray-800 p-6 rounded-lg shadow-md"
        >
            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                    <label className="block mb-1 font-semibold text-gray-700 dark:text-gray-300">
                        Event Name
                    </label>
                    <input
                        type="text"
                        value={data.name ?? ""}
                        // onChange={(e) => setName(e.target.value)}
                        onChange={(e) =>
                            setData({ ...data, name: e.target.value })
                        }
                        required
                        className="w-full px-3 py-2 rounded border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                    />
                </div>

                <div>
                    <label className="block mb-1 font-semibold text-gray-700 dark:text-gray-300">
                        Category
                    </label>
                    <input
                        type="text"
                        value={data.category ?? ""}
                        // onChange={(e) => setCategory(e.target.value)}
                        onChange={(e) =>
                            setData({ ...data, category: e.target.value })
                        }
                        required
                        className="w-full px-3 py-2 rounded border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                    />
                </div>
            </div>

            <div>
                <label className="block mb-1 font-semibold text-gray-700 dark:text-gray-300">
                    Description
                </label>
                <textarea
                    value={data.description ?? ""}
                    // onChange={(e) => setDescription(e.target.value)}
                    onChange={(e) =>
                        setData({ ...data, description: e.target.value })
                    }
                    required
                    rows={3}
                    className="w-full px-3 py-2 rounded border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                />
            </div>

            <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
                <div>
                    <label className="block mb-1 font-semibold text-gray-700 dark:text-gray-300">
                        Date & Time
                    </label>
                    <input
                        type="datetime-local"
                        defaultValue={
                            data.date
                                ? new Date(data.date).toISOString().slice(0, 16)
                                : ""
                        }
                        onChange={(e) => {
                            const localDate = new Date(e.target.value);
                            const rfcDate = localDate.toISOString();
                            setData({ ...data, date: rfcDate });
                        }}
                        required
                        className="w-full px-3 py-2 rounded border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                    />
                </div>

                <div>
                    <label className="block mb-1 font-semibold text-gray-700 dark:text-gray-300">
                        Venue
                    </label>
                    <input
                        type="text"
                        value={data.venue ?? ""}
                        // onChange={(e) => setVenue(e.target.value)}
                        onChange={(e) =>
                            setData({ ...data, venue: e.target.value })
                        }
                        required
                        className="w-full px-3 py-2 rounded border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                    />
                </div>

                <div>
                    <label className="block mb-1 font-semibold text-gray-700 dark:text-gray-300">
                        Price ($)
                    </label>
                    <input
                        type="number"
                        min="0"
                        step="0.01"
                        value={data.price ?? 0}
                        onChange={(e) =>
                            // setPrice(
                            //     e.target.value === ""
                            //         ? ""
                            //         : Number(e.target.value)
                            // )
                            setData({
                                ...data,
                                price:
                                    e.target.value === ""
                                        ? 0
                                        : Number(e.target.value),
                            })
                        }
                        required
                        className="w-full px-3 py-2 rounded border border-gray-300 dark:border-gray-600 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                    />
                </div>
            </div>

            <div>
                <label className="block mb-1 font-semibold text-gray-700 dark:text-gray-300">
                    Event Image
                </label>
                <input
                    type="file"
                    accept="image/*"
                    onChange={handleImageChange}
                    className="w-full text-gray-700 cursor-pointer dark:text-gray-300"
                />
            </div>
            <div>
                <div className="flex justify-between items-center mb-2">
                    <h2 className="text-xl font-semibold text-gray-800 dark:text-gray-100">
                        Translations
                    </h2>
                    <button
                        type="button"
                        onClick={handleAddTranslation}
                        className="px-3 py-1 cursor-pointer bg-blue-500 text-white rounded hover:bg-blue-600 transition"
                    >
                        + Add Translation
                    </button>
                </div>

                {data.translations?.length === 0 && (
                    <p className="text-gray-600 dark:text-gray-400 italic">
                        No translations added.
                    </p>
                )}

                {data.translations?.map((t, idx) => (
                    <div
                        key={idx}
                        className="mb-4 p-4 border border-gray-300 dark:border-gray-600 rounded-lg bg-gray-50 dark:bg-gray-700"
                    >
                        <div className="grid grid-cols-1 md:grid-cols-4 gap-4 mb-2">
                            <input
                                type="text"
                                placeholder="Language"
                                value={t.language ?? ""}
                                onChange={(e) =>
                                    handleTranslationChange(
                                        idx,
                                        "language",
                                        e.target.value
                                    )
                                }
                                className="px-3 py-2 rounded border bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                            />
                            <input
                                type="text"
                                placeholder="Name"
                                value={t.name ?? ""}
                                onChange={(e) =>
                                    handleTranslationChange(
                                        idx,
                                        "name",
                                        e.target.value
                                    )
                                }
                                className="px-3 py-2 rounded border bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                            />
                            <input
                                type="text"
                                placeholder="Venue"
                                value={t.venue ?? ""}
                                onChange={(e) =>
                                    handleTranslationChange(
                                        idx,
                                        "venue",
                                        e.target.value
                                    )
                                }
                                className="px-3 py-2 rounded border bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                            />
                            <button
                                type="button"
                                onClick={() => handleRemoveTranslation(idx)}
                                className="w-full px-3 py-2 bg-red-600 text-white rounded hover:bg-red-700 transition"
                            >
                                Remove
                            </button>
                        </div>
                        <textarea
                            placeholder="Description"
                            value={t.description ?? ""}
                            onChange={(e) =>
                                handleTranslationChange(
                                    idx,
                                    "description",
                                    e.target.value
                                )
                            }
                            className="w-full px-3 py-2 rounded border bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                        />
                    </div>
                ))}
            </div>

            <button
                type="submit"
                className="w-full py-3 cursor-pointer bg-green-600 text-white font-semibold rounded hover:bg-green-700 transition"
            >
                {submitLabel}
            </button>
        </form>
    );
};

export default EventForm;
