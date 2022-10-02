function Error() {
    return (
        <div>
            <h1>Error: {sessionStorage.getItem('error')}</h1>
            <button onClick={() => sessionStorage.removeItem('error')}>Clear Error</button>
        </div>
    );
}

export default Error;
