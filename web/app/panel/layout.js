"use client";

import Link from "next/link";
import { useEffect, useState } from "react";
import { IoPersonOutline } from "react-icons/io5";
import Addresses from "../components/addresses";
import Image from "next/image";

export default function PanelLayout({ children }) {
    const [user, setUser] = useState({ first_name: "", last_name: "" });
    const [addresses, setAddresses] = useState([]);

    useEffect(() => {
        fetch("http://localhost/api/profile/personal-info", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                authorization: "Bearer " + localStorage.getItem("access_token"),
            },
        })
            .then((response) => response.json())
            .then((res) => {
                setUser(res.data);

                // TODO: get addresses
                fetch("http://localhost/api/profile/addresses", {
                    headers: {
                        "Content-Type": "application/json",
                        authorization: "Bearer " + localStorage.getItem("access_token"),
                    },
                })
                    .then((resposne) => resposne.json())
                    .then((res) => setAddresses(res.data));
            });
    }, []);

    return (
        <>
            <header className="flex justify-between items-center h-28 border-b-2 shadow px-10 mb-8">
                <div>
                    <Link href="/">
                        <Image src="/images/Snappfood-Logo.png" width={100} height={100} alt="Snapp Food" />
                    </Link>
                </div>

                <div>
                    {addresses.length != 0 && <Addresses addresses={addresses} defaultAddress={user.default_address_id} />}
                </div>

                <div>
                    <Link href="/profile" className="flex items-center px-5 py-3 text-sm">
                        <IoPersonOutline fontWeight="bold" />
                        <span className="mx-2">
                            {user.first_name + " " + user.last_name}
                        </span>
                    </Link>
                </div>
            </header>

            <div className="container mx-auto px-6">{children}</div>
        </>
    );
}
