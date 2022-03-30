import Banner from './components/index/Banner';
import ToggleButton from './components/darkmode/ToggleButton';

const Index = () => { 
    
    return (
        <>
            <div className="absolute p-8 right-0">
                <ToggleButton />
            </div>
            <Banner/>
        </>
    )
}

export default Index;