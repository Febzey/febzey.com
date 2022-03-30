//@ts-ignore
import golang      from '../../images/golang.svg';
//@ts-ignore
import react       from '../../images/react.svg';
//@ts-ignore
import tailwindcss from '../../images/tailwindcss.svg';
//@ts-ignore
import typescript  from '../../images/typescript.svg';

const Banner = () => {


    return (
        <div className="h-screen dark:bg-zinc-900 w-full font-poppins flex items-center justify-center flex-col px-2 gap-4">

            <div className="flex flex-col gap-2 text-center absolute top-64">
                <h1 className="lg:text-8xl text-7xl font-bold text-indigo-500">Febzey</h1>
                <p className="dark:text-neutral-300 text-neutral-600 text-xl lg:text-2xl font-semibold">Full-Stack developer</p>
                <div className="p-2 text-center font-poppins">
                    <p className="text-neutral-800 dark:text-neutral-400 text-sm lg:text-md">Checkout my personal portfolio:
                        <a href="https://brayden.tech" className="pl-1 underline decoration-2 underline-offset-1 dark:text-indigo-400 text-indigo-600 duration-200 hover:text-indigo-600" target="_blank">brayden.tech</a></p>
                </div>
            </div>



            <div className="flex flex-col bottom-0 gap-3 mt-auto mb-8">
                <p className="text-center dark:text-neutral-400 text-sm">This website was made with</p>
                <div className="flex flex-row gap-4">
                    <img width="50" src={golang}></img>
                    <img width="50" src={react}></img>
                    <img width="50" src={typescript}></img>
                    <img width="50" src={tailwindcss}></img>
                </div>
            </div>


        </div>
    )
}

export default Banner;