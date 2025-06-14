export const formatTimestamp = (timestamp: string): string => {
    if (!timestamp || typeof timestamp !== "string") return "Invalid format";

    let timePart = "";

    if (timestamp.includes("T")) {
        timePart = timestamp.split("T")[1];
    } else if (/^\d{6}$/.test(timestamp)) {
        timePart = timestamp;
    } else if (/^\d{8}T\d{6}$/.test(timestamp)) {
        timePart = timestamp.substring(9);
    } else if (/^\d{2}:\d{2}(:\d{2})?$/.test(timestamp)) {
        timePart = timestamp.replace(/:/g, "");
    } else {
        const parsedDate = new Date(timestamp);
        if (!isNaN(parsedDate.getTime())) {
            return parsedDate.toTimeString().substring(0, 5);
        }
        return "Invalid format";
    }

    if (timePart.length < 4) return "Invalid format";

    const hours = timePart.substring(0, 2);
    const minutes = timePart.substring(2, 4);

    return `${hours}:${minutes}`;
};

export const formatDate = (dateString: string): string => {
    if (!dateString || typeof dateString !== "string") return "Invalid format";

    const date = new Date(dateString);
    if (isNaN(date.getTime())) return "Invalid format";

    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, "0");
    const day = String(date.getDate()).padStart(2, "0");

    return `${year}-${month}-${day}`;
};
