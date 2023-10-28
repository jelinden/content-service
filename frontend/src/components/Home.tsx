import image from './content.jpg'

const Home = () => {
    return (
        <>
            
            <div className="wrap" style={{marginLeft: '-40px', paddingLeft: '40px', paddingRight: '40px'}}>
                <h1 style={{paddingBottom: '30px', paddingTop: '30px'}}>Content service</h1>
                <div style={{display: 'inline-block', 
                    float: 'left', 
                    position: 'relative',
                    paddingTop: '1px',
                    fontWeight: 'bold',
                    fontSize: '26px',
                    color: '#3545A7'}}>
                    Content service is a service to hold your content and use it through an API.
                    This way it is easy to manage the content and for example language versions of the texts.
                </div>
                <img src={image} alt="content" style={{
                    width: '100%', 
                    display: 'inline-block',
                    opacity: '0.3',
                    position: 'absolute',
                    left: 0,
                    top: 0}}/>
            </div>
        </>

    )
}

export default Home