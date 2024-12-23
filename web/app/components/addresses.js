import { useState } from "react";
import { IoTrash } from "react-icons/io5";

const Addresses = ({ addresses, defaultAddress }) => {
    const [isOpen, setIsOpen] = useState(false);
    const currentAddress = addresses.find((addr) => addr.id == defaultAddress)

    const handleToggleModal = () => {
        setIsOpen((p) => !p);
    };

    const handleUpdateDefaultAddress = () => {
        // TODO: fix this
        console.log("update this to default address")
    }

    return (
        <div className="relative inline-block text-left w-100">
            <div>
                <button
                    onClick={handleToggleModal}
                    className="inline-flex justify-center w-full rounded-md
                    border border-gray-300 shadow-sm px-4 py-2 bg-white 
                    text-sm font-medium text-gray-700 hover:bg-gray-50 
                    focus:outline-none"
                >
                    {currentAddress.address}
                    <svg
                        className="ml-2 -mr-1 h-5 w-5"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                        aria-hidden="true"
                    >
                        <path
                            strokeLinecap="round"
                            strokeLinejoin="round"
                            strokeWidth="2"
                            d="M19 9l-7 7-7-7"
                        />
                    </svg>
                </button>
            </div>

            {isOpen && (
                <div className="fixed inset-0 flex
                        items-center justify-center
                        bg-black bg-opacity-50 text-right z-50">
                    <div className="bg-white rounded-lg
                            shadow-lg p-6 max-w-md
                            w-full relative">
                        <button
                            className="absolute top-2 right-2
                               text-gray-500 hover:text-gray-700"
                            onClick={handleToggleModal}
                        >
                            &#x2715; {/* Close button */}
                        </button>
                        <h1 className="text-center mb-4">انتخاب آدرس</h1>
                        <ul>
                            {addresses.map((addr) => {
                                return (
                                    <li
                                        className="border rounded-lg px-3 py-5 cursor-pointer flex justify-between content-between"
                                        key={addr.id}
                                        onClick={handleUpdateDefaultAddress}
                                    >
                                        <section className="flex items-center">
                                            <button className={`w-5 h-5 rounded-full border-2 border-blue-500 ml-3 ${defaultAddress === addr.id ? 'bg-blue-500' : ''}`}></button>
                                            <section>
                                                <h2>{addr.name}</h2>
                                                <p className="text-gray-500 text-sm">{addr.address}</p>
                                            </section>
                                        </section>
                                        {/* TODO: create delete address action */}
                                        <button className="text-red-700 text-2xl"><IoTrash /></button>
                                    </li>
                                )
                            })}

                        </ul>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Addresses;
