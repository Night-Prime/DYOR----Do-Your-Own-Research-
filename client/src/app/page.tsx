import Image from "next/image";

export default function Home() {
  return (
    <section className="w-[100vw] h-[100vh] bg-white text-black">
      <main className="w-full h-full flex flex-col justify-center items-center gap-10">
          <h1 className="text-9xl font-bold">D Y O R</h1>
          <h3 className="text-4xl">Do Your Own Research.</h3>
          <button className="bg-black text-white py-2 px-3 rounded-full cursor-pointer">Let's Build</button>
      </main>
    </section>
  );
}
