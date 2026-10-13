function ProfessionalCard({ professional }) {
  return (
    <div className="card">
      <div className="card-header">
        <h2>{professional.name}</h2>
      </div>
      <div className="card-body">
        <p className="card-address">📍 {professional.address}</p>
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
    </div>
  )
}

export default ProfessionalCard