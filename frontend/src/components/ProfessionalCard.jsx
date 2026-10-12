function ProfessionalCard({ professional }) {
    return (
        <div className="card">
            <h2>{professional.name}</h2>
            <p>{professional.address}</p>
            <a href={`tel:${professional.phone}`} className="call-button">
                Llamar
            </a>
            <a
                href={`https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(professional.address + ", Ushuaia")}`}
                target="_blank"
                rel="noopener noreferrer"
                className="map-button"
            >
                Cómo llegar
            </a>
        </div>
    )
}

export default ProfessionalCard