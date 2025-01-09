import Image from "next/image";
import {FaStar} from "react-icons/fa";
import Link from "next/link";

const Store = ({store}) => {
  return (
      <Link
          href={`/panel/stores/${store.slug}`}
          key={store.id}
          className="w-52 h-72 border ml-5 rounded-xl shadow overflow-hidden"
      >
          <section className="w-full h-1/3 relative mb-16">
              <Image
                  src={store?.store_type.image}
                  fill={true}
                  alt={store?.store_type.name}
              />
              <section
                  className="rounded-xl overflow-hidden absolute z-10 left-14"
                  style={{ bottom: "-50px" }}
              >
                  <Image
                      src={store.logo}
                      width={100}
                      height={100}
                      alt={store.name}
                  />
              </section>
          </section>
          <section className="flex flex-col items-center justify-center px-3">
              <span className="font-bold mb-2">{store.name}</span>
              <span className="inline-flex items-center text-xs mb-2">
                    <FaStar className="text-yellow-400 ml-2" /> {"۴.۳ (۳.۵۰۰)"}
                  </span>
              <span className="text-gray-700 text-sm">
                    {store.store_type.name}
                  </span>
          </section>
      </Link>
  );
}

export default Store;