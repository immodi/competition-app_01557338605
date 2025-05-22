const AdminDisclaimer: React.FC = () => {
    return (
        <div className="text-sm text-center text-yellow-700 dark:text-yellow-300 bg-yellow-100 dark:bg-yellow-900 border border-yellow-300 dark:border-yellow-700 rounded-lg p-3">
            In case you didn't read the github{" "}
            <span className="font-bold">README.md</span> carefully, the default
            admin credentials are:
            <div className="mt-4 flex flex-col">
                <div>
                    <span>Username:</span>{" "}
                    <code className="text-[1rem] font-bold font-mono">
                        admin
                    </code>
                </div>
                <div>
                    <span>Password:</span>{" "}
                    <code className="text-[1rem] font-bold font-mono">
                        admin
                    </code>
                </div>
            </div>
        </div>
    );
};

export default AdminDisclaimer;
