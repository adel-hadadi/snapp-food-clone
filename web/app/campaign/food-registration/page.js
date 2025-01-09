"use client";

import { useEffect, useState } from "react";

const Page = () => {
    const [provinces, setProvinces] = useState([]);
    const [cities, setCities] = useState([]);
    const [storeTypes, setStoreTypes] = useState([]);
    const [formData, setFormData] = useState({
        province_id: 0,
        city_id: 0,
        store_type_id: 0,
        name: "",
        latitude: 0,
        longitude: 0,
        address: "",
        logo: "",
        phone: "",
        manager_first_name: "",
        manager_last_name: ""
    });

    useEffect(() => {
        fetch("http://localhost/api/store-types")
            .then((res) => res.json())
            .then(({ data }) => setStoreTypes(data));
    }, []);

    const handleChange = (event) => {
        setFormData({ ...formData, [event.target.name]: event.target.value });
    };

    const handleSelecteProvince = (provinceID) => {
        setFormData({ ...formData, province_id: provinceID });

        fetch(`http://localhost/api/provinces/${provinceID}/cities`)
            .then((res) => res.json())
            .then(({ data }) => setCities(data));
    };

    const handleSelectStoreType = (storeTypeId) => {
        setFormData({ ...formData, store_type_id: storeTypeId });
    };

    const handleSelectCity = (cityId) => {
        setFormData({ ...formData, city_id: cityId });
    };

    const handleFormSubmit = (event) => {
        event.preventDefault();

        fetch("http://localhost/api/stores", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(formData),
        })
            .then((res) => res.json())
            .then((res) => console.log("response => ", res));
    };

    useEffect(() => {
        fetch("http://localhost/api/provinces")
            .then((res) => res.json())
            .then(({ data }) => setProvinces(data));
    }, []);

    return (
        <div className="flex h-lvh items-center justify-center">
            <div className="border shadow py-7 px-3 rounded">
                <form className="w-full max-w-lg" onSubmit={handleFormSubmit}>
                    <div className="w-full px-3 mb-6">
                        <label
                            className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                            htmlFor="first-name"
                        >
                            نام مالک
                        </label>
                        <input
                            className="appearance-none block w-full bg-gray-200 text-gray-700 border border-red-500 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white"
                            id="manager-first-name"
                            name="manager_first_name"
                            type="text"
                            onChange={handleChange}
                        />
                    </div>
                    <div className="w-full px-3 mb-6">
                        <label
                            className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                            htmlFor="last-name"
                        >
                            نام خانوادگی مالک
                        </label>
                        <input
                            className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                            id="last-name"
                            name="manager_last_name"
                            type="text"
                            onChange={handleChange}
                        />
                    </div>

                    <div className="w-full px-3 mb-6">
                        <label
                            className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                            htmlFor="resturant-name"
                        >
                            شماره تلفن مالک
                        </label>
                        <input
                            className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                            id="last-name"
                            name="phone"
                            type="text"
                            onChange={handleChange}
                        />
                    </div>

                    <div className="w-full px-3 mb-6">
                        <label
                            className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                            htmlFor="resturant-name"
                        >
                            نام فروشگاه
                        </label>
                        <input
                            className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                            id="last-name"
                            name="name"
                            type="text"
                            onChange={handleChange}
                        />
                    </div>

                    <div className="w-full  px-3 mb-6">
                        <label
                            className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                            htmlFor="state"
                        >
                            نوع فروشگاه
                        </label>
                        <div className="relative">
                            <select
                                className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                                id="state"
                                name="store_type_id"
                            >
                                {storeTypes.length > 0 &&
                                    storeTypes.map((type) => (
                                        <option
                                            key={type.id}
                                            onClick={() => handleSelectStoreType(type.id)}
                                        >
                                            {type.name}
                                        </option>
                                    ))}
                            </select>
                            <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                                <svg
                                    className="fill-current h-4 w-4"
                                    xmlns="http://www.w3.org/2000/svg"
                                    viewBox="0 0 20 20"
                                >
                                    <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
                                </svg>
                            </div>
                        </div>
                    </div>

                    <div className="flex flex-wrap -mx-3 mb-3">
                        <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
                            <label
                                className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                                htmlFor="state"
                            >
                                استان
                            </label>
                            <div className="relative">
                                <select
                                    className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                                    id="state"
                                    name="province_id"
                                >
                                    {provinces.length > 0 &&
                                        provinces.map((province) => (
                                            <option
                                                key={province.id}
                                                onClick={() => handleSelecteProvince(province.id)}
                                            >
                                                {province.name}
                                            </option>
                                        ))}
                                </select>
                                <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                                    <svg
                                        className="fill-current h-4 w-4"
                                        xmlns="http://www.w3.org/2000/svg"
                                        viewBox="0 0 20 20"
                                    >
                                        <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
                                    </svg>
                                </div>
                            </div>
                        </div>

                        <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
                            <label
                                className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                                htmlFor="state"
                            >
                                شهر
                            </label>
                            <div className="relative">
                                <select
                                    className="block appearance-none w-full bg-gray-200 border border-gray-200 text-gray-700 py-3 px-4 pr-8 rounded leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
                                    id="state"
                                    name="city_id"
                                >
                                    {cities.length > 0 &&
                                        cities.map((city) => (
                                            <option
                                                key={city.id}
                                                onClick={() => handleSelectCity(city.id)}
                                            >
                                                {city.name}
                                            </option>
                                        ))}
                                </select>
                                <div className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                                    <svg
                                        className="fill-current h-4 w-4"
                                        xmlns="http://www.w3.org/2000/svg"
                                        viewBox="0 0 20 20"
                                    >
                                        <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
                                    </svg>
                                </div>
                            </div>
                        </div>
                    </div>

                    <button
                        className="w-full py-2 bg-pink-600 rounded text-white"
                        type="submit"
                    >
                        ثبت
                    </button>
                </form>
            </div>
        </div>
    );
};

export default Page;
