"use client";

import {useParams} from "next/navigation";
import {useEffect, useState} from "react";
import Store from "@/app/components/store";

export default function Page({}) {
    const {slug} = useParams();
    const [stores, setStores] = useState([]);

    useEffect(() => {
        fetch(`http://localhost/api/panel/product-categories/${slug}/stores`, {
            headers: {
                authorization: "Bearer " + localStorage.getItem("access_token"),
            },
        })
            .then(res => res.json())
            .then(({data}) => setStores(data));
    }, [slug]);

    return (
        <div className="grid grid-cols-5">
            {stores.length > 0 && stores.map(store => <Store store={store} />)}
        </div>
    );
}
